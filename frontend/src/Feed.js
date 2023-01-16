import React from "react";
import "./feed.css";
import { Avatar } from "@mui/material";

function Feed({ username, image_url, caption, desc, user_image }) {
  return (
    <div className="feed">
      <img className="main__image" src={image_url} />
      <div className="feed__content">
        <h2 className="feed__text">{caption}</h2>
        <p className="desc__text"> {desc}</p>
        <div className="feed__header">
          <Avatar className="feed__avatar" src={user_image} />
          <text className="user__name">{username}</text>
        </div>
      </div>
    </div>
  );
}

export default Feed;
