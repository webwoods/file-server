# File Server with APIs README

This is a simple file server with APIs built in Go, allowing you to upload and retrieve images and videos.

## Getting Started

### Prerequisites

Make sure you have Go installed on your system. If not, you can download and install it from [the official Go website](https://golang.org/).

### Installation

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/Web-Woods/file-server.git
   ```

2. Navigate to the project directory:

   ```bash
   cd file-server
   ```

3. Build the project:

   ```bash
   go build
   ```

### Usage

1. Start the server:

   ```bash
   ./file-server
   ```

   The server will start running on `http://localhost:8080`.

2. You can access the server's APIs to upload and retrieve images and videos:

   - **Upload Image**: `POST /v1/api/upload/image`
   - **Get Image**: `GET /v1/api/get/image?filename=image.jpg`
   - **Get Images**: `GET /v1/api/get/images`
   - **Upload Video**: `POST /v1/api/upload/video`
   - **Get Video**: `GET /v1/api/get/video?filename=video.mp4`
   - **Get Videos**: `GET /v1/api/get/videos`

3. You can also use `curl` to check file uploads:

    ```
    curl -X POST -F "file=@/path/to/your/image.jpg" http://localhost:8080/v1/api/upload/image

    curl -X POST -F "file=@/path/to/your/video.mp4" http://localhost:8080/v1/api/upload/video
    ```
   
<!-- ## Security

To ensure that only authorized users have access to the server and its APIs, consider implementing authentication mechanisms such as API keys, JWT (JSON Web Tokens), OAuth, or OpenID Connect. -->

## Limitations

- The maximum file size for images is 20 MB, and for videos, it's 100 MB.
- Larger files exceeding the maximum size will not be uploaded.

<!-- ## Contributing

Contributions are welcome! If you encounter any issues or have suggestions for improvements, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. -->