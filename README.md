# OSU

**OpenSignedURL** is a simple file server with APIs built in Go, allowing you to upload files and get the experience of working with a CDN for free. 

*This is only suitable for development. (not suitable for production)*

## Getting Started

### Prerequisites

Make sure you have Go installed on your system. If not, you can download and install it from [the official Go website](https://golang.org/).  

Currently, OSU only supports MongoDB databases. Create a `.env` file in the root of the project folder and add your connection URI to connect to your database.

```
# .env
ds="mongodb://username:password@cluster/options"
```

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

2. To upload a file to the server, you must first request a pre-signed upload URL which will be the temporary file upload path.

   ```
   curl -X POST http://localhost:8080/v1/api/presigned/upload

   # sample response
   {"url":"/static/1711949164033836200/images/1711949164034402900.jpg","expiration":"2024-04-01T11:56:04.0344029+05:30"}
   ```

3. You can then use this URL to upload the file. Note that,
   - if the URL is invalid
   - if the folderId doesn't match
   - if the URL is expired

   then the file will not be uploaded.

   ```
   $ curl -X PUT -T img6.jpg http://localhost:8080/static/1711949164033836200/images/1711949164034402900.jpg

   # invalid URL response
   Failed to create file
   
   # expired URL response
   Presigned URL is expired

   # succesful upload response
   File 1711949164034402900.jpg uploaded successfully
   ```
   
<!-- ## Security

To ensure that only authorized users have access to the server and its APIs, consider implementing authentication mechanisms such as API keys, JWT (JSON Web Tokens), OAuth, or OpenID Connect. -->

## Limitations

   - Authentication features are still under development.
   - Maximum file size threshold not implemented.
   - Issuing API key is still under development
 
<!-- - The maximum file size for images is 20 MB, and for videos, it's 100 MB.
- Larger files exceeding the maximum size will not be uploaded. -->

<!-- ## Contributing

Contributions are welcome! If you encounter any issues or have suggestions for improvements, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. -->