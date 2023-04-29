import { get } from "./api/rest";

let response = async () => {
    let data = get();

    data.then( (result) => {
        console.log(result)
    });
};
response();