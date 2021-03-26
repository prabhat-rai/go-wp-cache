package web

// PLAN
	// 2 URL only
		// load_wp.com/call_wp/blog
		// load_wp.com/call_wp/tips
	// load data from .env file for WP URL & construct different URL parameter
	// call the APIs sequentially
	// connect to Redis and save the data in Redis
		// Redis Keys will also be in env variable

	// Future enhancement
		// Call the Wordpress APIs parallely [tags/categories/post-list]
		// run this service periodically on its own, without invoking.


func main() {

}