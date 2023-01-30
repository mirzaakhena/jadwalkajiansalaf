import to from "await-to-js";
import axios from "axios";
import {BASE_URL} from "./url";

export const orderSubmitted = async (payload) => {

    const url = `${BASE_URL}/admin`

    const [err, res] = await to(axios.post(url, payload).catch((err) => Promise.reject(err)))

    if (err) {
        return Promise.reject(err.response.data)
    }

    return Promise.resolve(res.data)
}