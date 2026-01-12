package main

import (
	"github/Tshewang2022/social/internal/auth"
	"github/Tshewang2022/social/internal/db"
	"github/Tshewang2022/social/internal/env"
	"github/Tshewang2022/social/internal/mailer"
	"github/Tshewang2022/social/internal/store"
	"time"

	"go.uber.org/zap"
)

const version = "0.0.9"

func main() {

	//	@title			GopherSocial API
	//	@description	API for GopherSocial, a social network for gohpers
	//	@termsOfService	http://swagger.io/terms/

	//	@contact.name	API Support
	//	@contact.url	http://www.swagger.io/support
	//	@contact.email	support@swagger.io

	//	@license.name	Apache 2.0
	//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

	//	@BasePath					/v1
	//
	//	@securityDefinitions.apikey	ApiKeyAuth
	//	@in							header
	//	@name						Authorization
	//	@description
	cfg := config{
		addr:        env.GetString("ADDR", ":8080"),
		frontendURL: env.GetString("FRONTEND_URL", "http://localhost:4000"),
		apiURL:      env.GetString("EXTERNAL_URL", "localhost:8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social? sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdelTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		mail: mailConfig{
			exp:       time.Hour * 24 * 3, // user have 3 days to accept the invitations
			fromEmail: env.GetString("FROM_EMAIL", ""),
			sendGrid: sendGridConfig{
				apiKey: env.GetString("SENDGRID_API_KEY", ""),
			},
		},

		auth: authConfig{
			basic: basicConfig{
				user: env.GetString("AUTH_BASIC_USER", "admin"),
				pass: env.GetString("AUTH_BASIC_USER", "admin"),
			},
			token: tokenConfig{
				secret: env.GetString("AUTH_TOKEN_SECRET", "example"),
				exp:    time.Hour * 24 * 3, // 3days
			},
		},
		env: env.GetString("ENV", "development"),
	}

	//logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdelTime,
	)

	if err != nil {
		logger.Panic(err)
	}

	defer db.Close()
	logger.Info("database connection pool established")

	store := store.NewStorage(db)

	mailer := mailer.NewSendGrid(cfg.mail.sendGrid.apiKey, cfg.mail.fromEmail)
	jwtAuthenticator := auth.NewJTWAuthenticator(cfg.auth.token.secret, cfg.auth.token.iss, cfg.auth.token.iss)

	app := &application{
		config:        cfg,
		store:         store,
		logger:        logger,
		mailer:        mailer,
		authenticator: jwtAuthenticator,
	}
	mux := app.mount()
	logger.Fatal(app.run(mux))
}
