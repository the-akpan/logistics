import express from "express";
import mongoose from "mongoose";
import path from "path";

const app = express();
const MONGODB_URI =
  process.env.MONGODB_URI || "mongodb://127.0.0.1:27017/logistics";
const PORT = process.env.PORT || 3000;

const __dirname = new URL(".", import.meta.url).pathname;

mongoose.set("strictQuery", false);
mongoose
  .connect(MONGODB_URI)
  .then(() => {
    const time = new Date();
    console.log(`${time.toISOString()}: MongoDB connected`);
  })
  .catch((err) => console.error(err));

// Register react
app.use(express.static("client/dist"));
app.get("/*", (req, res) => {
  res.sendFile(path.join(__dirname, "client", "dist", "index.html"));
});

app.listen(PORT, () => {
  console.log("Running on port " + PORT);
});
