# Go micro-service in ~30 minutes

This is a Go micro-service written from scratch.

It shows how to use [net/http](https://godoc.org/net/http), and how to structure a Go project.

It relies on Go 1.11 Beta 2 and the upcoming "Go modules" (formerly known as "vgo") support.

Dependency injection is used to insert a logger instance into the handler.

You can also notice how the test is constructed in order to provide testing for the handler.

A Docker container is available, thanks to the Dockerfile. It shows how to construct such containers. 

## How to use

Because this project uses go modules, as long as you are using Go 1.11 Beta 2+ or Go 1.10 with vgo support,
you should be ok.

Clone this anywhere in your computer and create a project in your editor. I'm using [GoLand IDE](https://jetbrains.com/go) in order to work
on the project during the presentation as well as have support for go modules.

The bundled, self-signed, certificates are bound to either ` dev.localhost:8080 ` or ` docker.localhost:8080 `. I obviously
do not recommend using these in production.

## Presentation link

I created this as part of the presentation at [GopherCon UK 18](https://www.gophercon.co.uk/).

The link for the video will be updated here when the presentation is out.

## References 

### Structuring Go applications

In order to learn how to approach package design in Go, you can read the following resources:

- [Style guideline for Go packages - JBD](https://rakyll.org/style-packages/)
- [Standard Package Layout - Ben Johnson](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1#.ds38va3pp)
- [Go best practices, six years in - Peter Bourgon](https://peter.bourgon.org/go-best-practices-2016/#repository-structure)

Once done, this article will help you understand the [Design Philosophy On Packaging by William Kennedy](https://www.ardanlabs.com/blog/2017/02/design-philosophy-on-packaging.html).

### Exposing Go applications to the Internet

[This article](https://blog.cloudflare.com/exposing-go-on-the-internet/) describes how you can start approaching  

## Thank you

I would like to thank you [William "Bill" Kennedy](https://twitter.com/goinggodotnet) for the inspiration he provided on
getting me to do this talk.


## License

This project is under the MIT license. Please see the [LICENSE](LICENSE) file for more details.
