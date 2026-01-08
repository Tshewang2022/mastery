package db

import (
	"context"
	"fmt"
	"github/Tshewang2022/social/internal/store"
	"log"
	"math/rand"
)

// database seeding
var usernames = []string{
	"alex_jones", "samuel99", "john_doe", "linda_k", "mike_dev", "sarah01", "daniel_s",
	"emma_w", "chris_lee", "olivia_m", "noah_t", "lucas88",
	"ethan_r", "ava_smith", "liam_k", "mia_j", "jackson_p", "amelia_c", "henry_b", "grace_l",
	"leo_dev", "nina_92", "ryan_h", "bella_x", "kevin_m", "zoe_r", "aaron_f", "sophia_k",
	"adam_007", "lucy_w", "victor_n", "harper_t", "james_d", "isabella_p", "mark_dev",
	"chloe_s", "brandon_l", "ella_m", "steven_h", "paul_r", "victoria_j",
	"timothy_k", "rebecca_s", "frank_b", "peter_w", "nancy_t", "andrew_c",
	"katie_m", "robert_x",
}

var titles = []string{
	"Getting Started with Go",
	"Understanding REST APIs",
	"Building Scalable Backend Systems",
	"Introduction to Databases",
	"How Authentication Works",
	"Designing Clean Architecture",
	"Working with WebSockets",
	"API Security Best Practices",
	"Handling Errors Gracefully",
	"Optimizing Application Performance",
	"Microservices Explained",
	"Deploying Applications to the Cloud",
	"Version Control with Git",
	"Writing Maintainable Code",
	"Debugging Like a Pro",
	"Using Docker for Development",
	"Testing Backend Applications",
	"System Design Basics",
	"Common Backend Mistakes",
	"From Monolith to Microservices",
}

var contents = []string{
	"This post explains the basic concepts needed to get started and understand the fundamentals clearly.",
	"In this article, we explore common challenges developers face and how to solve them effectively.",
	"This content walks through practical examples and best practices used in real-world applications.",
	"Learn how different components work together to build scalable and maintainable systems.",
	"This post covers essential tips that can improve performance and code quality.",
	"A simple explanation of complex ideas to help beginners gain confidence.",
	"This article focuses on clean design principles and why they matter in modern software development.",
	"Discover techniques that help reduce bugs and improve long-term maintainability.",
	"This content highlights common mistakes and how to avoid them during development.",
	"A practical guide that demonstrates step-by-step implementation details.",
	"This post provides insight into architectural decisions and trade-offs.",
	"An overview of tools and technologies commonly used in backend development.",
	"This article explains how to structure projects for better readability and scalability.",
	"Learn strategies for debugging issues efficiently in production environments.",
	"This content introduces testing approaches that ensure application reliability.",
	"A discussion on security considerations every developer should be aware of.",
	"This post explains how to handle data correctly and avoid common pitfalls.",
	"An introduction to deployment workflows and automation basics.",
	"This article shares lessons learned from real-world development experience.",
	"A concise explanation of best practices for building robust applications.",
}

func Seed(store store.Storage) {
	ctx := context.Background()

	users := generateUsers(100)

	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("Error creating users:", err)
			return
		}
	}

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	comments := generateComments(500, users, posts)

	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	log.Println("seeding is completed")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			Password: "1234",
		}
	}
	return users

}

var tags = []string{
	"golang",
	"backend",
	"api",
	"database",
	"programming",
	"webdev",
	"software",
	"architecture",
	"security",
	"performance",
	"testing",
	"docker",
	"microservices",
	"cloud",
	"devops",
	"rest",
	"scalability",
	"debugging",
	"bestpractices",
	"coding",
}

var comments = []string{
	"Great post, very helpful!",
	"This explanation made things much clearer.",
	"Thanks for sharing this information.",
	"I learned something new today.",
	"Well written and easy to understand.",
	"This answered a lot of my questions.",
	"Nice breakdown of the topic.",
	"Looking forward to more posts like this.",
	"This was exactly what I needed.",
	"Clear and concise explanation.",
	"Good insights, thanks!",
	"I appreciate the practical examples.",
	"This topic is explained really well.",
	"Helpful content, keep it up!",
	"Interesting perspective on this.",
	"Thanks for taking the time to write this.",
	"This is useful for beginners.",
	"Solid explanation overall.",
	"I enjoyed reading this post.",
	"Very informative article.",
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: titles[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}
	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)

	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}
	return cms
}
