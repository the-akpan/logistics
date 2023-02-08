import express from "express";
import mongoose from "mongoose";

const app = express();
const MONGODB_URI =
  process.env.MONGODB_URI || "mongodb://127.0.0.1:27017/logistics";
const PORT = process.env.PORT || 3000;

mongoose.set("strictQuery", false);
mongoose
  .connect(MONGODB_URI)
  .then(() => {
    const time = new Date();
    console.log(`${time.toISOString()}: MongoDB connected`);
  })
  .catch((err) => console.error(err));

app.listen(PORT, () => {
  console.log("Running on port " + PORT);
});
