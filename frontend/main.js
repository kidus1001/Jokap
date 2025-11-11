document.addEventListener ("DOMContentLoaded", () => {
    const loadButton = document.querySelector (".all-jokes-section button");
    const loadInput = document.getElementById ("loadId");

    const getJokeBtn = document.querySelector(".search-section button");
    const idOfJoke = document.getElementById ("jokeId");

    loadButton.addEventListener("click", async ()=> {
        const numberOfJokes = parseInt(loadInput.value);
        if (isNaN(numberOfJokes) || numberOfJokes <= 0) {
            alert("Please enter a valid number of jokes!");
            return;
        }
        const allJokesDiv = document.getElementById("allJokes");
        try {
            console.log("Number of jokes to load:", numberOfJokes);

            allJokesDiv.innerHTML = `<p>Loading ${numberOfJokes} jokes.</p>`

            const response = await fetch (`http://localhost:8080/jokes?num=${numberOfJokes}`, {
                method: "GET",
                headers: {
                    "Content-Type": "application/json",
                },
            });

            const data = await response.json();
            if (response.ok && data.success) {
                allJokesDiv.innerHTML = data.data.map((joke) => `<p><b>${joke.id}</b>: ${joke.content}</p>`).join("");
            } else {
                allJokesDiv.innerHTML = `<p>Error: ${data.error || "Unknown error"}.</p>`;
            }
        } catch (error){
            console.log("Error: ", error);
            allJokesDiv.innerHTML = `<p>Failed to fetch jokes.</p>`;
        }
    })


    getJokeBtn.addEventListener ("click", async () => {
        const id = parseInt(idOfJoke.value);
        if (isNaN(id) || id <= 0) {
            alert("Please enter a valid number of jokes!");
            return;
        }

        const singleJokeDiv = document.getElementById("singleJoke")
        try {
            console.log("Loading joke number: ", id);

            singleJokeDiv.innerHTML = `Loading joke num - ${id}`;

            const response = await fetch (`http://localhost:8080/joke/${id-1}`);

            const data = await response.json()

            if (response.ok) {
                const joke = data.data
                singleJokeDiv.innerHTML = `<p><b>${joke.id}</b>: ${joke.content}</p>`;
            } else {
                singleJokeDiv.innerHTML = `<p>Error: ${data.error || "Unknown error"}.</p>`;
            }


        } catch (error){
            console.log("Error: ", error);
            singleJokeDiv.innerHTML = `<p>Failed to fetch jokes.</p>`;
        }
    })
})