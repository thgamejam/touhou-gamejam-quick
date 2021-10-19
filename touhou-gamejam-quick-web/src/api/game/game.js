import axios from "axios";

const getGame = id => axios.get(`/api/v1/game/${id}`).then(res => res.data);

export {
  getGame
}