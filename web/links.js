async function HandleDelete(id){
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

async function HandleCopy(id) {
    let host = window.location.protocol + "//" + window.location.host;
    await navigator.clipboard.writeText(host + "/" + id);
}
