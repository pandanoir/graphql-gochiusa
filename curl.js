const request = require('superagent');

(async() => {
    let query =`
query {
    shop(name: "ラビットハウス") {
        name
    }
}`;
    console.log( (await request.post('http://localhost:8080').send(query)).text );
    query = `
query {
    shop(name: "ラビットハウス") {
        members {
            name, age
        }
    }
}`;
    console.log( (await request.post('http://localhost:8080').send(query)).text );
})();
