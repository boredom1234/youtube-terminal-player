# 🎬 youtube-terminal-player 💻

> Watch Youtube Videos in your Terminal!

This project lets you play YouTube videos directly in your terminal using ASCII characters. It's a fun project that demonstrates various programming concepts like downloading files, processing images, and manipulating the terminal output.

## ✨ Features

*   Downloads YouTube videos.
*   Extracts frames from the video.
*   Converts frames to ASCII art.
*   Plays the ASCII frames in the terminal, giving the illusion of a video.
*   Plays audio in the background for a complete experience.
*   Displays a progress bar for downloads and frame conversion.

## 🚀 Installation

1.  **Install FFmpeg:** This project uses FFmpeg for video and audio processing. Make sure you have it installed and added to your system's PATH.
    *   You can download it from [https://ffmpeg.org/download.html](https://ffmpeg.org/download.html).
2.  **Install Go:**  If you don't have Go installed, download and install it from [https://golang.org/](https://golang.org/).
3.  **Get the Project:** Clone this repository to your computer:

    ```bash
    git clone https://github.com/boredom1234/youtube-terminal-player.git
    cd youtube-terminal-player
    ```
4.  **Install Dependencies**

```bash
go get github.com/kkdai/youtube/v2
go get golang.org/x/image/draw
```

## ▶️ Usage

1.  **Run the program:**

```bash
go run main.go
```

1.  **Enter the YouTube video URL:**
    The program will prompt you to enter the URL of the YouTube video you want to watch. 

1.  **Enjoy the show!**

## ⚙️ Configuration

You can customize the video resolution and frame rate by changing these constants in the `main.go` file:

```go
const (
	FPS          = 60  // Frames per second
	Width        = 120 // ASCII frame width
	Height       = 67  // ASCII frame height
)
```

**Note:** Higher resolution and frame rate will require more processing power.

## 🤝 Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## 🙏 Acknowledgements

*   This project utilizes the [kkdai/youtube](https://github.com/kkdai/youtube) package for downloading YouTube videos.
*   The ASCII art conversion is inspired by various online resources.

## 😄 Enjoy!

Have fun watching YouTube videos in a whole new way! 
