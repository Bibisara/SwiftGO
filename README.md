# Welcome to SwiftGO!

SwiftGO is a web application framework that allows developers to build high-performance applications using the Go programming language. It is designed to be fast, secure, and easy to use.

## Getting Started

To get started with SwiftGO, you first need to install Go. You can download and install the latest version of Go from the official website: [GoLang](https://golang.org/dl/)

Once you have Go installed, you can install SwiftGO using the following command:
```
go get github.com/swiftgo-framework/swiftgo
```
> This will install SwiftGO and its dependencies.

Creating a new project

To create a new SwiftGO project, run the following command:
```
swiftgo new <project-name>
```
> This will create a new project with the specified name in your current directory.

## Routing

SwiftGO uses a simple and flexible routing system that allows you to map URLs to controller actions. Here's an example of how to define a route:
```ruby
app := swiftgo.New()

app.Get("/hello/:name", func(c *swiftgo.Context) {
    name := c.Param("name")
    c.String(http.StatusOK, "Hello %s", name)
})

app.Run(":8080")
```
This route will match URLs that start with `/hello/` followed by a name parameter. When the route is matched, the function defined as the second argument will be executed. The function has access to the request context, which can be used to retrieve information about the request.

## Controllers

Controllers are responsible for handling requests and returning responses. Here's an example of how to define a controller:

```ruby
type HelloController struct {}

func (c *HelloController) SayHello(ctx *swiftgo.Context) {
    name := ctx.Param("name")
    ctx.String(http.StatusOK, "Hello %s", name)
}
```
This controller defines a single action called `SayHello`. The action takes a request context as an argument and returns a response.

## Middleware

Middleware allows you to define custom behavior that will be executed before or after a request is processed. Here's an example of how to define middleware:

```ruby
func Logger() swiftgo.HandlerFunc {
    return func(c *swiftgo.Context) {
        log.Printf("[%s] %s", c.Request.Method, c.Request.URL.Path)
        c.Next()
    }
}

app := swiftgo.New()

app.Use(Logger())

app.Get("/hello/:name", func(c *swiftgo.Context) {
    name := c.Param("name")
    c.String(http.StatusOK, "Hello %s", name)
})

app.Run(":8080")
```
This example defines a middleware function called `Logger`. The function logs the HTTP method and URL of the request, then calls the `Next` function to continue processing the request.

## Conclusion

SwiftGO is a fast, secure, and easy-to-use web application framework for the Go programming language. With its simple and flexible routing system, powerful controllers, and customizable middleware, SwiftGO makes it easy to build high-performance web applications. Try it out today!
