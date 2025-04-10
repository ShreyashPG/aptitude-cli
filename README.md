# Aptitude Test CLI 🧠🖥️

A command-line based aptitude test application built with Go, designed to test and score users within a set timer. Questions are loaded from a JSON file and categorized by difficulty levels.

---

## ✨ Features

- ⏱️ Customizable timer for the whole test
- 🎯 Choose difficulty level: easy, medium, hard, or random
- ✅ Each question is unique (no repeats during a session)
- 📊 Instant scoring at the end
- 🗂️ Questions loaded from a JSON file
- 🧪 Clean and interactive CLI experience

---

## 📁 Project Structure

```
aptitude-cli/
│
├── main.go                # Entry point
├── questions.json         # All aptitude questions
├── go.mod                 # Go module file
└── utils/
    ├── file.go            # Load and filter questions
    └── timer.go           # Countdown timer logic
```

---

## 🛠️ Getting Started

### 1. Clone the repo
```bash
git clone https://github.com/ShreyashPG/aptitude-cli.git
cd aptitude-cli
```

### 2. Build the CLI
```bash
go build -o aptitude-test
```

### 3. Run the test
```bash
./aptitude-test --time=60 --level=easy
```

**Flags:**
- `--time` = total test time in seconds
- `--level` = question difficulty (`easy`, `medium`, `hard`, `random`)

---

## 📝 questions.json Format

Each question must follow this structure:

```json
{
  "question": "What is 2 + 2?",
  "options": ["1", "2", "3", "4"],
  "answer": "4",
  "difficulty": "easy"
}
```

- Store all 90 questions in `questions.json`
- Difficulty should be one of: `easy`, `medium`, `hard`

---

## 📌 TODO

- [ ] Add per-question timer (optional)
- [ ] Save scores to file
- [ ] Review incorrect answers at the end
- [ ] Export score summary

---

## 👨‍💻 Author

**Shreyash Ghanekar**  
[GitHub: ShreyashPG](https://github.com/ShreyashPG)

---

## 🧾 License

This project is licensed under the MIT License. Feel free to fork, improve, and share!
