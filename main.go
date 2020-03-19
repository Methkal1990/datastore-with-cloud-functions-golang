package main


import (
	"context"
	"cloud.google.com/go/datastore"
	"time"
	"fmt"
	
)

type Post struct {
    Title string
    PublishedAt  time.Time
}

func main(){
	Dts()
}


func Dts() {

	ctx := context.Background()
	
	client, _ := datastore.NewClient(ctx, "wave26-sellbytel-methkal")
	
	opt, err := client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		
		keys := []*datastore.Key{
			datastore.NameKey("Post", "post1", nil),
			datastore.NameKey("Post", "post2", nil),
			datastore.NameKey("Post", "post4", nil),
			datastore.NameKey("Post", "post5", nil),
			datastore.NameKey("Post", "post6", nil),
			datastore.NameKey("Post", "post7", nil),
			datastore.NameKey("Post", "post8", nil),
			datastore.NameKey("Post", "post9", nil),
			datastore.NameKey("Post", "post10", nil),
			datastore.NameKey("Post", "post11", nil),
			datastore.NameKey("Post", "post12", nil),
			datastore.NameKey("Post", "post13", nil),
			datastore.NameKey("Post", "post14", nil),
			datastore.NameKey("Post", "post15", nil),
		}

		posts := []*Post{
			{Title: "Post 1", PublishedAt: time.Now()},
			{Title: "Post 2", PublishedAt: time.Now()},
			{Title: "Post 3", PublishedAt: time.Now()},
			{Title: "Post 4", PublishedAt: time.Now()},
			{Title: "Post 5", PublishedAt: time.Now()},
			{Title: "Post 6", PublishedAt: time.Now()},
			{Title: "Post 7", PublishedAt: time.Now()},
			{Title: "Post 8", PublishedAt: time.Now()},
			{Title: "Post 9", PublishedAt: time.Now()},
			{Title: "Post 10", PublishedAt: time.Now()},
			{Title: "Post 11", PublishedAt: time.Now()},
			{Title: "Post 12", PublishedAt: time.Now()},
			{Title: "Post 13", PublishedAt: time.Now()},
			{Title: "Post 14", PublishedAt: time.Now()},
			{Title: "Post 15", PublishedAt: time.Now()},
		}
		_, err := tx.PutMulti(keys, posts)
		if err != nil {
			return err
		}

		return nil
	})
	fmt.Println(err)
	fmt.Println(opt)
}
