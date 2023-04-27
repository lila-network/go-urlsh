async function HandleSubmit() {
    document.getElementById("submit").active = false
    let slug = document.getElementById("linkname").value
    let url = document.getElementById("link").value
    let description = document.getElementById("description").value
    let body = {
        "id" : slug,
        "url" : url,
        "description" : description
    }

    let response = await fetch("/api/v1/links", {
        credentials: "include",
        body: JSON.stringify(body),
        mode: "same-origin",
        method: "POST"
    });

    if (!response.ok) {
        document.getElementById("dialog-heading").textContent = "Error"
        document.getElementById("dialog-text").textContent = "The following error occured during the request: " + response.statusText
        document.getElementById('dialog-info').showModal()

        document.getElementById("submit").active = true
    }
}

async function HandleChange() {
    console.log("HandleChange")
    let buttonactive = true
    if (document.getElementById("link").value === "")
    {
        buttonactive = false
    }
    document.getElementById("submit").active = buttonactive
}