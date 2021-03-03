<h1 align="center">
    Quiz app
</h1>

<p align="center">
  <a href="#-project">Project</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#rocket-technologies">Technologies</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#information_source-how-to-run">How to run</a>&nbsp;&nbsp;&nbsp;

</p>

<br>


## 💻 Project

This is my first application written in Go. Just a simple quiz app.

## :rocket: Technologies

This app was made using the following technologies:
- [Go][go]
## :information_source: How To Run

To clone and run this application, you'll need [Git](https://git-scm.com) + [Go][go] installed on your computer.

### Clone and Run

<br>

```bash
# Clone this repository
$ git clone https://github.com/lucastssb/quiz-app.git

# Go into the repository
$ cd quiz-app

# Compile and run the app
$ build . && ./quiz-app

# The default time limit for each question is 3 seconds
# The default csv problems file is problems.csv


```

### Use custom settings

To use a different csv file use the flag "-csv"

Example:

```bash
$ go build . && ./quiz-app -csv=questions.csv
```
<br>

To set a different time limit use the flag "-limit"

Example:

```bash
$ go build . && ./quiz-app -limit=10
```

<br>

Made with ♥ by Lucas Barbosa :wave: [Get in touch!](https://www.linkedin.com/in/lucas-barbosa-60b56416b/)

[go]: https://golang.org/

