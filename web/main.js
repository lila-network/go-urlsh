// Link overview

async function HandleLinkIndexDelete(id){
    let response = await fetch("/api/v1/links/" + id, {
        credentials: "include",
        mode: "same-origin",
        method: "DELETE"
    });
    if (!response.ok) {
        console.log("error deleting " + id + ": " + response.statusText)
    }
    document.location = "/admin/"
}

async function HandleLinkIndexCopy(id) {
    let host = window.location.protocol + "//" + window.location.host;
    await navigator.clipboard.writeText(host + "/" + id);
}


// Link Add

async function HandleLinkAddSubmit() {
    await LinkAction("add")
}

// Link edit

async function HandleLinkEditSubmit() {
    await LinkAction("edit")
}

async function LinkAction(action){
    document.getElementById("submit").active = false
    let slug = document.getElementById("linkname").value
    let url = document.getElementById("link").value
    let description = document.getElementById("description").value

    let method, endpoint = ""
    let body;
    switch(action) {
        case "add":
            method = "POST"
            endpoint = "/api/v1/links/"
            body = {
                "id" : slug,
                "url" : url,
                "description" : description
            }
            break;
        case "edit":
            method = "PUT"
            endpoint = "/api/v1/links/" + slug
            body = {
                "url" : url,
                "description" : description
            }
            break;
        default:
            return
    }



    let response = await fetch(endpoint, {
        credentials: "include",
        body: JSON.stringify(body),
        mode: "same-origin",
        method: method
    });

    if (!response.ok) {
        document.getElementById("dialog-heading").textContent = "Error"
        document.getElementById("dialog-text").textContent = "The following error occured during the request: " + response.statusText
        document.getElementById('dialog-info').showModal()

        document.getElementById("submit").active = true
    }
    document.location = "/admin/"
}

async function HandleLinkFieldChange() {
    console.log("HandleChange")
    let buttonactive = true
    if (document.getElementById("link").value === "")
    {
        buttonactive = false
    }
    document.getElementById("submit").active = buttonactive
}

// ApiKey Add

async function HandleApiKeyNewSubmit() {
    let button = document.getElementById("submit")
    let description = document.getElementById("description")
    button.active = false
    button.setAttribute("aria-busy", "true")

    let body = {
        "description" : description
    }

    let response = await fetch("/api/v1/apikeys", {
        credentials: "include",
        body: JSON.stringify(body),
        mode: "same-origin",
        method: "POST"
    });

    if (response.ok) {
        let data = await response.json()
        document.getElementById("dialog-heading").textContent = "New API-Key"
        document.getElementById("dialog-text").textContent = "Here is your new API Key. Copy it NOW, it won't be shown again."
        document.getElementById("dialog-apikey").textContent = data.key
        document.getElementById('dialog-info').showModal()
    }
}

async function HandleApiKeyModalClose() {
    let modal = document.getElementById('dialog-info');
    modal.close()
}

// General

function Logout() {
    document.cookie = 'gourlsh_auth=; Max-Age=-1; path=/; domain=' + location.hostname;
    document.location = "/admin/login"
}
