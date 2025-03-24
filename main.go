package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	// Path to your Firebase service account key (downloaded from Firebase console)
	sa := option.WithCredentialsFile("ai-teaching-a31ac-firebase-adminsdk-fbsvc-3328f096cd.json")

	client, err := firestore.NewClient(ctx, "ai-teaching-a31ac", sa)
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}
	defer client.Close()

	fmt.Println("Connected to Firestore!")

	// Create schema
	createSchema(ctx, client)
}

func createSchema(ctx context.Context, client *firestore.Client) {
	// User schema
	_, _, err := client.Collection("users").Add(ctx, map[string]interface{}{
		"name":         "Alice Johnson",
		"email":        "alice@email.com",
		"role":         "student",
		"profile_pic":  "https://link-to-pic",
		"classroom_id": "class123",
		"streak":       5,
		"achievements": []string{"badge1", "badge2"},
		"created_at":   firestore.ServerTimestamp,
		"last_login":   firestore.ServerTimestamp,
	})
	if err != nil {
		log.Fatalf("Failed to create user schema: %v", err)
	}
	fmt.Println("User schema created!")

	// Classroom schema
	_, _, err = client.Collection("classrooms").Add(ctx, map[string]interface{}{
		"name":       "Physics 101",
		"teacher_id": "teacher456",
		"students":   []string{"userId123", "userId456"},
		"created_at": firestore.ServerTimestamp,
		"updated_at": firestore.ServerTimestamp,
	})
	if err != nil {
		log.Fatalf("Failed to create classroom schema: %v", err)
	}
	fmt.Println("Classroom schema created!")

	// Assignment schema
	_, _, err = client.Collection("assignments").Add(ctx, map[string]interface{}{
		"title":        "Newton's Laws Problem Set",
		"description":  "Solve the provided questions.",
		"classroom_id": "class123",
		"due_date":     "2025-03-25T23:59:00Z",
		"total_marks":  100,
		"submissions": map[string]interface{}{
			"userId123": map[string]interface{}{
				"file_url":     "https://drive-link-to-file",
				"submitted_at": firestore.ServerTimestamp,
				"score":        85,
				"feedback":     "Good understanding, but needs more detail.",
				"graded_by":    "AI",
			},
		},
	})
	if err != nil {
		log.Fatalf("Failed to create assignment schema: %v", err)
	}
	fmt.Println("Assignment schema created!")

	// ✅ Fixed Progress schema
	_, _, err = client.Collection("progress").Add(ctx, map[string]interface{}{
		"user_id":         "userId123",
		"classroom_id":    "class123",
		"topics_mastered": []string{"Kinematics", "Dynamics"},
		"weak_areas":      []string{"Thermodynamics"},
		"quiz_scores": []map[string]interface{}{
			{
				"quiz_id": "quiz456",
				"score":   78,
			},
			{
				"quiz_id": "quiz789",
				"score":   92,
			},
		},
		"learning_trends": map[string]interface{}{
			"last_week":  "improving",
			"last_month": "consistent",
		},
		"last_updated": firestore.ServerTimestamp,  // ✅ ServerTimestamp outside the array
	})
	if err != nil {
		log.Fatalf("Failed to create progress schema: %v", err)
	}
	fmt.Println("Progress schema created!")

	// Gamification schema
	_, _, err = client.Collection("gamification").Add(ctx, map[string]interface{}{
		"user_id":          "userId123",
		"points":           320,
		"badges":           []string{"Persistent Learner", "Quiz Master"},
		"streak_days":      7,
		"leaderboard_rank": 12,
	})
	if err != nil {
		log.Fatalf("Failed to create gamification schema: %v", err)
	}
	fmt.Println("Gamification schema created!")

	// ✅ Fixed AI Insights schema
	_, _, err = client.Collection("ai_insights").Add(ctx, map[string]interface{}{
		"user_id": "userId123",
		"personalized_recommendations": []string{
			"Review Newton's 2nd Law.",
			"Watch video on Conservation of Energy.",
		},
		"learning_path": map[string]interface{}{
			"current_level":        "Intermediate",
			"suggested_next_topic": "Work and Energy",
		},
		"generated_quizzes": []map[string]interface{}{
			{
				"quiz_id":    "quiz789",
				"title":      "Forces & Motion Quiz",
				"difficulty": "Medium",
			},
		},
		"last_updated": firestore.ServerTimestamp,  // ✅ ServerTimestamp outside the array
	})
	if err != nil {
		log.Fatalf("Failed to create AI insights schema: %v", err)
	}
	fmt.Println("AI Insights schema created!")
}
