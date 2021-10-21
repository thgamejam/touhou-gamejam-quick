import axios from "axios";

const getGame = id => axios.get(`/api/v1/game/${id}`).then(res => res.data);

const getGameDownloadURL = id => axios.get(`/api/v1/game/download/${id}`).then(res => res.data);

const getGameList = page => axios.get(`/api/v1/game/list/${page}`).then(res => res.data);

export {
  getGame,
  getGameDownloadURL,
  getGameList,
}