function getProjects(){
    let xhr = new XMLHttpRequest();
    xhr.open("GET", "/projects", false);
    xhr.setRequestHeader("Authorization", "Bearer " + localStorage.getItem("token"));
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.send();

    let responseObj = JSON.parse(xhr.responseText);

    if (!responseObj.status){
        alert(responseObj.message);
    }

    if (responseObj.projects != null)
    responseObj.projects.forEach((project) => {
       renderProject(project.id, project.name, project.tasks);
    });
}

let jwt = localStorage.getItem("token");

if (!jwt){
    renderLoginForm();
}else{
    getProjects();
    renderProjectAddBtn()
}



