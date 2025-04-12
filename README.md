
# Aptitude CLI Application

This is a Go-based CLI aptitude test application that provides various test options including time-based tests, random question tests, and difficulty-based tests (Easy, Medium, Hard). It loads questions from a JSON file (`Aptitude.json`) and allows the user to take quizzes interactively.

## Features

- **Time-based Aptitude Test**: User can set a time limit and answer questions within the given time frame.
- **Random Questions**: Randomized questions from the available pool.
- **Difficulty-based Questions**: User can choose questions based on difficulty level: Easy, Medium, or Hard.
- **Interactive CLI**: The user interacts with the application via terminal input.
- **Scoring**: The app tracks and displays the user's score at the end of the test.

## Installation

1. Clone the repository or download the code.
2. Make sure you have [Go](https://golang.org/doc/install) installed on your machine.
3. Place the `Aptitude.json` file (containing the questions data) in the same directory as the Go code.

## Usage

### Step 1: Run the Docker container

Build the Docker image for the app:

```bash
docker build -t aptitude-cli .
```

Run the Docker container with the interactive flag:

```bash
docker run --rm -it aptitude-cli
```

### Step 2: Select a Test Option

After running the app, you will be presented with the following options:

1) Time Limit Aptitude Test
2) Solve Questions Randomly
3) Solve Easy Questions
4) Solve Medium Questions
5) Solve Hard Questions

### Step 3: Answer the Questions

Once you select an option, follow the on-screen instructions to take the test. You will be prompted to answer each question, and the app will provide feedback on whether your answer is correct or wrong.

### Step 4: View Results

At the end of the test, the application will display your score.

## Requirements

- Go 1.21 or higher
- Docker (for running in containerized environment)

## License

MIT License
