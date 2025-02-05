document.getElementById('open_btn').addEventListener('click', function() {
    document.getElementById('sidebar').classList.toggle('open-sidebar');
})

// handle requests from button event

// computer buttons event
let computerBtn = document.getElementById("computerGet")
computerBtn.addEventListener("click", function() {    
    fetch("http://localhost:8080/computer/", {
        headers: {
            "Content-Type": "application/json",
        },
    })
    .then((response) => response.json())
    .then((data) => {
        received = JSON.stringify(data)
        if (data.error) {
            console.log(data.message)
        } else {
            console.log(received)
        }
    })
    .catch((error) => {
        output.innerHTML += "<br><br>Error: " + error
    })
});