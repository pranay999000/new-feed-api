import React, { useState, useEffect } from "react";
import "./App.css";
import icon from "./images/small_icon.png";
import feed2Icon from "./images/feed2.png";
import dany from "./images/dany_cropped.jpg";
import handmaid from "./images/handmaid.jpg";
import harry from "./images/harry.png";
import hermione from "./images/hermione.jpg";
import indiana from "./images/indiana_jones.jpg";
import jack from "./images/jack_sparrow.jpeg";
import leia from "./images/leia.jpg";
import mia from "./images/mia_cropped.jpg";
import norton from "./images/norton.jpg";
import seb from "./images/seb.jpg";
import tooper from "./images/trooper.jpg";
import sharp from "./images/sharp3.jpg";
import logout from "./images/logout.png";
import Feed from "./Feed";
import Profile from "./Profile";
import { Avatar, Button, Input, Modal } from "@mui/material";
import { makeStyles } from "@mui/styles";
import Login from "./Login";

import { ReactSession } from "react-client-session";

const map = new Map();
map.set("dany", dany);
map.set("handmaid", handmaid);
map.set("harry", harry);
map.set("hermione", hermione);
map.set("indiana", indiana);
map.set("jack", jack);
map.set("leia", leia);
map.set("mia", mia);
map.set("norton", norton);
map.set("seb", seb);
map.set("tooper", tooper);
map.set("sharp", sharp);

function getModalStyles() {
  const top = 50;
  const left = 50;

  return {
    top: `${top}%`,
    left: `${left}%`,
    transform: `translate(-${top}%, -${left}%)`,
  };
}

const useStyles = makeStyles((theme) => ({
  paper: {
    position: "absolute",
    width: 400,
    backgroundColor: "white",
    border: "2px solid #000",
    boxShadow: theme.boxShadow,
    padding: "18px",
    borderRadius: 8,
  },
}));

function App() {
  ReactSession.setStoreType("localStorage");

  const [feeds, setFeeds] = useState([]);
  const [open, setOpen] = useState(false);
  const [signUpText, setSignUpText] = useState(true);
  const classes = useStyles();
  const [modalStyle] = useState(getModalStyles);

  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [image, setImage] = useState("");
  const [password, setPassword] = useState("");

  const [user, setUser] = useState(null);
  const [home, setHome] = useState(0);

  useEffect(() => {
    if (ReactSession.get("token")) {
      setHome(0);
      setUser({
        name: ReactSession.get("name"),
        image: ReactSession.get("image"),
        email: ReactSession.get("email"),
      });
    } else {
      setHome(1);
    }
  }, []);

  const feedList = [
    {
      username: "Pranay",
      image_url: indiana,
      caption: "American epic space opera multimedia",
      desc: "Star Wars is an American epic space opera multimedia franchise created by George Lucas, which began with the eponymous 1977 film and quickly became a worldwide pop-culture phenomenon. The franchise has been expanded into various films and other media, including television series, video games, novels, comic books, theme park attractions, and themed areas, comprising an all-encompassing fictional universe. Star Wars is one of the highest-grossing media franchises of all time. StarWars is an American epic space opera multimedia franchise created byGeorge Lucas, which began with the eponymous 1977 film and quicklybecame a worldwide pop-culture phenomenon. The franchise has beenexpanded into various films and other media, including televisionseries, video games, novels, comic books, theme park attractions, andthemed areas, comprising an all-encompassing fictional universe. StarWars is one of the highest-grossing media franchises of all time. StarWars is an American epic space opera multimedia franchise created byGeorge Lucas, which began with the eponymous 1977 film and quicklybecame a worldwide pop-culture phenomenon. The franchise has beenexpanded into various films and other media, including televisionseries, video games, novels, comic books, theme park attractions, andthemed areas, comprising an all-encompassing fictional universe. StarWars is one of the highest-grossing media franchises of all time. StarWars is an American epic space opera multimedia franchise created byGeorge Lucas, which began with the eponymous 1977 film and quicklybecame a worldwide pop-culture phenomenon. The franchise has beenexpanded into various films and other media, including televisionseries, video games, novels, comic books, theme park attractions, andthemed areas, comprising an all-encompassing fictional universe. StarWars is one of the highest-grossing media franchises of all time.",
    },
    {
      username: "Pranay",
      image_url: dany,
      caption: "American epic space opera multimedia",
      desc: "Star Wars is an American epic space opera multimedia franchise created by George Lucas, which began with the eponymous 1977 film and quickly became a worldwide pop-culture phenomenon. The franchise has been expanded into various films and other media, including television series, video games, novels, comic books, theme park attractions, and themed areas, comprising an all-encompassing fictional universe. Star Wars is one of the highest-grossing media franchises of all time. Star Wars is an American epic space opera multimedia franchise created by George Lucas, which began with the eponymous 1977 film and quickly became a worldwide pop-culture phenomenon. The franchise has been expanded into various films and other media, including television series, video games, novels, comic books, theme park attractions, and themed areas, comprising an all-encompassing fictional universe. Star Wars is one of the highest-grossing media franchises of all time. Star Wars is an American epic space opera multimedia franchise created by George Lucas, which began with the eponymous 1977 film and quickly became a worldwide pop-culture phenomenon. The franchise has been expanded into various films and other media, including television series, video games, novels, comic books, theme park attractions, and themed areas, comprising an all-encompassing fictional universe. Star Wars is one of the highest-grossing media franchises of all time. Star Wars is an American epic space opera multimedia franchise created by George Lucas, which began with the eponymous 1977 film and quickly became a worldwide pop-culture phenomenon. The franchise has been expanded into various films and other media, including television series, video games, novels, comic books, theme park attractions, and themed areas, comprising an all-encompassing fictional universe. Star Wars is one of the highest-grossing media franchises of all time.",
    },
  ];

  useEffect(() => {
    setFeeds(feedList.map((f) => f));
  }, []);

  const signUp = (event) => {
    event.preventDefault();
  };

  const renderProfile = (event) => {
    event.preventDefault();
    setHome(2);
  };

  const handlePage = (event) => {
    event.preventDefault();
    if (user != null) {
      setHome(0);
    } else {
      setHome(1);
    }
  };

  const handleLogout = (event) => {
    event.preventDefault();
    setUser(null);
    setHome(1);
  };

  return (
    <div className="app">
      <Modal open={open} onClose={() => setOpen(false)}>
        <div style={modalStyle} className={classes.paper}>
          <form className="app__signup">
            <center>
              <img className="icon__signup" src={icon} />
            </center>
            <Input
              placeholder="Username"
              type="text"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
            />
            <Input
              placeholder="Email"
              type="text"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
            <Input
              placeholder="Password"
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
            <text className="sign__up2" onClick={() => setOpen(true)}>
              Sign up
            </text>
          </form>
        </div>
      </Modal>
      <div className="app__header">
        <img
          className="app__headerImage"
          src={feed2Icon}
          alt="Logo"
          onClick={handlePage}
        />
        <div className="app__userDetails">
          {user != null ? (
            <div className="app__header__top__right">
              <div className="app__user" onClick={renderProfile}>
                <h4 className="app__userName">{user.name}</h4>
                <Avatar className="feed__avatar" src={map.get(user.image)} />
              </div>
              <img className="logout" onClick={handleLogout} src={logout} />
            </div>
          ) : signUpText ? (
            <text className="sign__up" onClick={() => setSignUpText(false)}>
              Sign in
            </text>
          ) : (
            <text className="sign__up" onClick={() => setSignUpText(true)}>
              Sign up
            </text>
          )}
        </div>
      </div>
 

      {home === 0 ? (
        feeds.map((f) => (
          <Feed
            username={f.username}
            image_url={f.image_url}
            caption={f.caption}
            desc={f.desc}
          />
        ))
      ) : home === 1 ? (
        <Login instance={signUpText} />
      ) : (
        <Profile />
      )}
    </div>
  );
}

export default App;
