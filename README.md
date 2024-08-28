# QUIZ GAME

This repository contains the source code for a fun and engaging quiz game written in Go. Test your knowledge across various topics and see how well you score!

## Purpose

The purpose of this project was to learn the basics of Go Language. I am planning to learn Backend Development in Go. This was a short project I did to get familiar with the Go Lang.

## **Features**

* **JSON-based questions:** Define your quiz questions and answers in a JSON file for easy maintenance and flexibility. You can easily add your own questions by following below structure:

  ```json
  [
    {
      "text": "What is the capital of France?",
      "options": [
        "London",
        "Paris",
        "Berlin",
        "Madrid"
      ],
      "answer": 2 // Index of the correct answer (starts from 1)
    },
    // ... more questions ...
  ]
  ```
* **Interactive gameplay:** The game provides a visually appealing and interactive experience with timed text effects.
* **Age verification:** Ensures the quiz is appropriate for a wider audience.
* **Detailed feedback:** Provides feedback on each answer, including correct options.
* **Scoring system:** Tracks your performance throughout the quiz and displays your final score.
* **Encouragement messages:** Motivates players based on their score.

## Building & Running

1. Clone the repository.
2. Install Go from [https://golang.org/doc/install](https://golang.org/doc/install).
3. Make sure you have a JSON file named "questions.json" in the project directory with your quiz questions structured as shown above (replace the example content with your own).
4. Build the program using command `go build -o bin/app app.go`
5. Then run `bin/app` to run the program

I have already build for windows and saved it as app-windows.exe inside the bin folder. If you are on windows, you can simply run it to experience the program without needing to install Go.

You can also just enter `go run app.go` inside the project directory to run the program without saving the build. This will make you skip step 4 & 5.

## **License**

This project is licensed under the MIT License (see LICENSE file for details).

# **Enjoy the Quiz!**
