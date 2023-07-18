# Imagine

_🌄✨ A magical image transformation server ✨🌄_

<p align="center">
    <img src="https://media3.giphy.com/media/26BRKOtNgpAqlnxf2/giphy.gif?cid=ecf05e47n0poatxkof4fxyp70qpbqbq70xop2h8n9c6fa86x&ep=v1_gifs_search&rid=giphy.gif&ct=g" alt="Imagine GIF">
</p>

> 🌟💡 **Disclaimer**: Please be aware that this project was created solely for educational purposes. 📚 While I've strived for code quality, I cannot guarantee its compatibility with your specific project. Therefore, I kindly request that you exercise caution when integrating this code into your own work. My main objective was to grasp the fundamentals of image transformation in Golang and expose them through an API endpoint. It may not be the most optimized code I've written, but it definitely fulfills its purpose. 🎯

## ✨🌟 Amazing Features 🌟✨

**🖼️ Image Magic ✨**

- ✅ Rotate images
- 📏 Resize images
- 📐 Scale images
  (more coming soon)

**🌈 Filter Fun ✨**

- 🌻 Sepia filter
- 🌑 Negative filter
- 🌕 Grayscale filter
  (more coming soon)

**🔍 Meta Magic ✨**

- Get Meta Data

(And many more...)

## 🚀 Installation

To get started with go-imagine, make sure you have the following requirements:

- Go >= 1.20.4

📥 Clone the repository:

```bash
  git clone https://github.com/Muhammed-Rajab/Imagine.git
```

📂 Navigate to the project directory:

```bash
  cd Imagine/
```

⬇️ Download the necessary Go modules:

```bash
  go mod download
```

🏃‍♀️ Run the application:

```bash
  go run .
```

This will start the server at http://localhost:4000/

## API Reference

#### 🔥 Transform Image

```http
  GET /api/images/transform
```

| Query parameter | Type      | Description                                                                                |
| :-------------- | :-------- | :----------------------------------------------------------------------------------------- |
| `rotate`        | `integer` | **optional**. Angle to which you want your image to be rotated (in degrees) (Default: `0`) |
| `scale`         | `integer` | **optional**. Scale to which you want your image to be scaled (Default: `1`)               |
| `resize`        | `string`  | **optional**. New size of your image. eg: `resize=256x256` (Default: `None`)               |

| Form data | Type       | Description                                                                 |
| :-------- | :--------- | :-------------------------------------------------------------------------- |
| `image`   | `jpeg/png` | **required**. The image file to which the transformation must be applied to |

#### ✨ Apply filter

```http
  GET /api/images/filter
```

| Query parameter | Type     | Description                                         |
| :-------------- | :------- | :-------------------------------------------------- |
| `name`          | `string` | **required**. Filter to apply on the provided image |

Possible filter names: `sepia`, `grayscale`, `negative`

| Form data | Type       | Description                                                                 |
| :-------- | :--------- | :-------------------------------------------------------------------------- |
| `image`   | `jpeg/png` | **required**. The image file to which the transformation must be applied to |

#### 🖼️ Get image metadata

```http
  GET /api/images/meta
```

| Form data | Type       | Description                                                                 |
| :-------- | :--------- | :-------------------------------------------------------------------------- |
| `image`   | `jpeg/png` | **required**. The image file to which the transformation must be applied to |
