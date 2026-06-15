# ASCII Studio

![ASCII Studio](web/static/images/logo.png) <!-- Update logo path if applicable -->

**ASCII Studio** is a powerful, responsive web application that transforms standard text into beautiful, customizable ASCII art. Built with a Go backend and a vanilla HTML/CSS/JS frontend, it provides digital artists, developers, and retro enthusiasts an instant, zero-lag environment to generate ASCII banners.

## 🚀 Features

* **Real-time Generation**: See your ASCII art generated instantly as you type.
* **Custom Banners**: Choose from classic styles such as `Standard`, `Shadow`, and `Thinkertoy`.
* **Color Customization**: Add a splash of color to your terminal outputs or digital graphics using a full hex color picker.
* **Alignment Controls**: Align your ASCII art to the Left, Center, or Right.
* **Project Saving**: Save your configuration (text, banner, color, and alignment) and get a unique shareable URL.
* **Export Options**: 
  * Instantly copy the generated ASCII to your clipboard.
  * Download the art as a `.txt` file for use in code comments, readmes, or terminal scripts.
* **Fully Responsive**: A mobile-first design ensures the workspace is perfectly usable across desktops, laptops, tablets, and smartphones.

## 🛠️ Technology Stack

* **Backend**: Go (Golang) - Handles the API routing, ASCII conversion logic, and serves static files.
* **Frontend**: Vanilla HTML5, CSS3, and JavaScript - No heavy frameworks, ensuring blazing fast load times and a lightweight footprint.
* **Storage**: Local JSON file storage for saving/loading projects.
* **Containerization**: Docker & Docker Compose ready.

## 📦 Getting Started

### Prerequisites

* [Go](https://golang.org/doc/install) 1.18 or higher (if running locally)
* [Docker](https://docs.docker.com/get-docker/) & Docker Compose (if running via containers)

### Running Locally (Without Docker)

1. Clone the repository:
   ```bash
   git clone https://github.com/Manjaws1/ASCII-Studio.git
   cd ASCII-Studio
   ```

2. Start the Go server:
   ```bash
   go run cmd/server/main.go
   ```

3. Open your browser and navigate to:
   ```
   http://localhost:8087
   ```

### Running with Docker

1. Ensure the Docker daemon is running.
2. Build and start the container using Docker Compose:
   ```bash
   docker-compose up --build
   ```
3. Open your browser and navigate to:
   ```
   http://localhost:8087
   ```

## 📂 Project Structure

```text
.
├── cmd/
│   └── server/
│       └── main.go         # Application entry point and HTTP handlers
├── data/                   # Directory where saved projects are stored (JSON)
├── internal/
│   ├── alignment/          # Logic for text alignment
│   ├── ascii/              # Core ASCII generation and banner parsing logic
│   ├── color/              # Logic for applying colors to the output
│   ├── renderer/           # Orchestrates ASCII generation, color, and alignment
│   └── storage/            # Handles saving and loading projects
├── web/
│   ├── static/             # CSS stylesheets and JavaScript files
│   └── templates/          # HTML templates (index.html)
├── Dockerfile              # Docker configuration for the Go app
├── docker-compose.yml      # Docker Compose configuration
└── README.md               # Project documentation
```

## 🤝 Contributing

Contributions are welcome! Feel free to open an issue or submit a Pull Request if you have ideas for new features, banner styles, or UI improvements.

