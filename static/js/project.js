let projectsList = document.getElementById("projects-list");

let format_date = "yyyy-mm-dd";

function renderProject(id, title, tasks) {
    let projectBlock = document.createElement("div");
    projectBlock.classList.add("project");
    projectBlock.setAttribute("id-project", id);

    let projectHeaderBlock = document.createElement("div");
    projectHeaderBlock.classList.add("project__header");
    let projectCalendarIcon = document.createElement("i");
    projectCalendarIcon.classList.add("far", "fa-calendar-alt", "project__header-calendar");
    let projectTitle = document.createElement("h4");
    projectTitle.classList.add("project__header-title");
    projectTitle.innerText = title;
    let projectEditIcon = document.createElement("i");
    projectEditIcon.classList.add("fas", "fa-pencil-alt", "project__header-edit");
    let projectDeleteIcon = document.createElement("i");
    projectDeleteIcon.classList.add("far", "fa-trash-alt", "project__header-delete");

    // Project Edit Name
    let projectHeaderEditInput = document.createElement("input");
    projectHeaderEditInput.classList.add("project__header-input");
    projectHeaderEditInput.value = title;
    let projectHeaderEditBtn = document.createElement("button");
    projectHeaderEditBtn.classList.add("project__header-btn");
    projectHeaderEditBtn.innerText = "Confirm";
    // -------------------------------------------------------------------------------

    // Project task block
    let projectTaskAddBlock = document.createElement("div");
    projectTaskAddBlock.classList.add("project__task-add");
    let projectPlusIcon = document.createElement("i");
    projectPlusIcon.classList.add("fas", "fa-plus");

    // Add deadline input
    let projectTaskDeadLineInput = document.createElement("input");
    projectTaskDeadLineInput.placeholder = "Click here to select deadline...";
    projectTaskDeadLineInput.classList.add("project__task-add-deadline");
    // -----------------------------------

    let projectTaskInput = document.createElement("input");
    projectTaskInput.classList.add("project__task-add-title");
    projectTaskInput.placeholder = "Start typing here to create a task...";
    let projectTaskButton = document.createElement("button");
    projectTaskButton.innerText = "Add Task";
    projectTaskButton.setAttribute("id-project", id);
    // ---------------------------------------------

    projectTaskAddBlock.appendChild(projectPlusIcon);

    // Add deadline Input
    projectTaskAddBlock.appendChild(projectTaskDeadLineInput);
    // -----------------------------------------------

    projectTaskAddBlock.appendChild(projectTaskInput);
    projectTaskAddBlock.appendChild(projectTaskButton);

    projectHeaderBlock.appendChild(projectCalendarIcon);
    projectHeaderBlock.appendChild(projectTitle);

    // Project Edit Name
    projectHeaderBlock.appendChild(projectHeaderEditInput);
    projectHeaderBlock.appendChild(projectHeaderEditBtn);
    // ---------------------------------------------------------

    projectHeaderBlock.appendChild(projectEditIcon);
    projectHeaderBlock.appendChild(projectDeleteIcon);


    projectBlock.appendChild(projectHeaderBlock);
    projectBlock.appendChild(projectTaskAddBlock);
    projectBlock.appendChild(renderTasksBlock(id, tasks));

    projectsList.insertBefore(projectBlock, projectsList.firstChild);

    //  Add deadline input
    let dp = new Datepicker(projectTaskDeadLineInput, {
        autohide: true,
        format: format_date,
        minDate: new Date(),
        title: "Select task deadline",
    });
    //  ----------------------

    projectTaskButton.addEventListener('click', event => {
        let idProject = Number(event.target.getAttribute("id-project"));
        let input = document.querySelector("div[id-project=\"" + idProject + "\"] .project__task-add .project__task-add-title");
        let inputDeadLine = document.querySelector("div[id-project=\"" + idProject + "\"] .project__task-add .project__task-add-deadline");

        let text = input.value;
        let deadline = inputDeadLine.value;

        let projectTasksBlock = document.querySelector("div[id-project=\"" + idProject + "\"] .project__tasks");


        if (validateTask(text, deadline)) {
            //  checked and order

            // Send request for add task
            let xhr = new XMLHttpRequest();
            xhr.open("POST", "/task/create", false);
            xhr.setRequestHeader("Content-type", "application/json");
            xhr.setRequestHeader("Authorization", "Bearer " + localStorage.getItem("token"));

            let requestObj = {};
            requestObj.text = text;
            requestObj.deadline = deadline;
            requestObj.id_project = idProject;

            xhr.send(JSON.stringify(requestObj));

            let responseObj = JSON.parse(xhr.responseText);

            if (!responseObj.status) {
                alert(responseObj.message);
                return
            }

            let task = responseObj.task;
            // -----------------------------------------------------

            projectTasksBlock.appendChild(renderTasks(task.id, task.text, task.checked, task.deadline, task.priority, task.id_project));

            input.value = '';
            inputDeadLine.value = '';
            dp.setDate({clear: true});
        } else {
            alert("Input correct task text or deadline");
        }
    });

    projectTaskInput.addEventListener("keypress", e => {
        if (e.key === "Enter") {
            projectTaskButton.click();
        }
    });

    projectDeleteIcon.addEventListener('click', () => {
        if (!confirm("You sure to want delete this project?")) return;


        let xhr = new XMLHttpRequest();
        xhr.open("DELETE", "/project/delete", false);
        xhr.setRequestHeader("Content-Type", "application/json");
        xhr.setRequestHeader("Authorization", "Bearer " + localStorage.getItem("token"));

        let project = {}; project.id = id;

        xhr.send(JSON.stringify(project));

        let responseObj = JSON.parse(xhr.responseText);

        if (!responseObj.status) {
            alert(responseObj.message);
            return
        }


        projectBlock.remove();
    });

    projectEditIcon.addEventListener('click', () => {
        if (projectTitle.style.display !== "none") {
            projectTitle.style.display = "none";
            projectHeaderEditInput.style.display = "block";
            projectHeaderEditBtn.style.display = "block";
            projectHeaderEditInput.focus();
        }
    });

    projectHeaderEditInput.addEventListener('keypress', event => {
        if (event.key === "Enter") {
            projectHeaderEditBtn.click();
        }
    });

    projectHeaderEditBtn.addEventListener('click', () => {
        let temp_title = projectHeaderEditInput.value;

        if (temp_title.length < 4) {
            alert("Project title must be at least 4 characters");
        } else if (temp_title === title) {
            projectTitle.style.display = "block";
            projectHeaderEditInput.style.display = "none";
            projectHeaderEditBtn.style.display = "none";
        } else {
            let xhr = new XMLHttpRequest();
            xhr.open("PUT", "/project/edit", false);
            xhr.setRequestHeader("Content-Type", "application/json");
            xhr.setRequestHeader("Authorization", "Bearer " + localStorage.getItem("token"));

            let project = {};
            project.name = temp_title;
            project.id = id;

            xhr.send(JSON.stringify(project));

            let responseObj = JSON.parse(xhr.responseText);

            if (!responseObj.status) {
                alert(responseObj.message);
                return
            }

            title = temp_title;

            projectTitle.innerText = title;
            projectTitle.style.display = "block";
            projectHeaderEditInput.style.display = "none";
            projectHeaderEditBtn.style.display = "none";
        }
    });


}

function renderTasksBlock(idProject, tasks) {
    let projectTasksBlock = document.createElement("div");
    projectTasksBlock.classList.add("project__tasks");
    let redLine = document.createElement("div");
    redLine.classList.add("red-double-line");
    let blueLine = document.createElement("div");
    blueLine.classList.add("blue-line");
    projectTasksBlock.appendChild(redLine);
    projectTasksBlock.appendChild(blueLine);

    if (tasks != null)
        tasks.forEach((element) => {
            projectTasksBlock.appendChild(renderTasks(element.id, element.text, element.checked, element.deadline, element.priority, idProject));
        });


    return projectTasksBlock
}

function renderTasks(id, text, checked, deadline, order, idProject) {
    let taskBlock = document.createElement("div");
    taskBlock.classList.add("task");
    taskBlock.setAttribute("order", order);
    let taskCheckBox = document.createElement("input");
    taskCheckBox.type = "checkbox";
    taskCheckBox.classList.add("task__check");
    taskCheckBox.checked = checked;
    let taskText = document.createElement("p");
    taskText.classList.add("task__text");
    taskText.innerText = text;

    // Add task deadline
    let taskDeadline = document.createElement("p");
    taskDeadline.innerHTML = "<b>Deadline: " + deadline + "</b>";
    taskText.appendChild(taskDeadline);
    // -------------------------


    // Task edit block
    let taskEditBlock = document.createElement("div");
    taskEditBlock.classList.add("task__edit-block");
    let taskEditTextarea = document.createElement("textarea");
    taskEditTextarea.classList.add("task__edit-textarea");
    let taskEditDeadlineInput = document.createElement("input");
    taskEditDeadlineInput.classList.add("task__edit-deadline");
    taskEditDeadlineInput.placeholder = "Select new deadline";
    let taskEditConfirmButton = document.createElement("button");
    taskEditConfirmButton.innerText = "Confirm";
    let taskEditCancelButton = document.createElement("button");
    taskEditCancelButton.innerText = "Cancel";

    taskEditTextarea.value = text;
    taskEditDeadlineInput.value = deadline;

    let dp = new Datepicker(taskEditDeadlineInput, {
        autohide: true,
        format: format_date,
        minDate: new Date(),
        title: "Select task deadline",
    });
    dp.setDate(deadline);

    taskEditBlock.appendChild(taskEditDeadlineInput);
    taskEditBlock.appendChild(taskEditTextarea);
    taskEditBlock.appendChild(taskEditConfirmButton);
    taskEditBlock.appendChild(taskEditCancelButton);

    taskEditCancelButton.addEventListener('click', () => {
        if (taskEditBlock.style.display !== "none") {
            taskEditTextarea.value = text;
            dp.setDate(deadline);
            taskEditBlock.style.display = "none";
            taskText.style.display = "block";
        }
    });

    taskEditConfirmButton.addEventListener('click', () => {
        let temp_text = taskEditTextarea.value;
        let temp_deadline = taskEditDeadlineInput.value;

        if (validateTask(temp_text, temp_deadline)) {

            let xhr = new XMLHttpRequest();
            xhr.open("PUT", "/task/edit", false);
            xhr.setRequestHeader("Content-Type", "application/json");
            xhr.setRequestHeader("Authorization", "Bearer " + localStorage.getItem("token"));

            let task = {};
            task.id = id;
            task.id_project = idProject;
            task.text = temp_text;
            task.deadline = temp_deadline;

            xhr.send(JSON.stringify(task));

            let responseObj = JSON.parse(xhr.responseText);

            if (!responseObj.status) {
                alert(responseObj.message);
                return
            }

            text = temp_text;
            deadline = temp_deadline;
            taskText.innerText = text;
            taskDeadline.innerHTML = "<b>Deadline: " + deadline + "</b>";
            taskText.appendChild(taskDeadline);
            taskEditBlock.style.display = "none";
            taskText.style.display = "block";
        } else {
            alert("Input correct task text or deadline");
        }
    });
    //---------------------------------

    let taskSettingsBlock = document.createElement("div");
    taskSettingsBlock.classList.add("task__settings");
    let taskArrowUp = document.createElement("i");
    taskArrowUp.classList.add("fas", "fa-angle-up", "task__settings-item");
    let taskArrowDown = document.createElement("i");
    taskArrowDown.classList.add("fas", "fa-angle-down", "task__settings-item");
    let taskEditIcon = document.createElement("i");
    taskEditIcon.classList.add("fas", "fa-pencil-alt", "task__settings-item");
    let taskDeleteIcon = document.createElement("i");
    taskDeleteIcon.classList.add("fas", "fa-trash-alt", "task__settings-item");

    taskBlock.appendChild(taskCheckBox);
    taskBlock.appendChild(taskText);
    taskBlock.appendChild(taskEditBlock);
    taskBlock.appendChild(taskSettingsBlock);
    taskSettingsBlock.appendChild(taskArrowUp);
    taskSettingsBlock.appendChild(taskArrowDown);
    taskSettingsBlock.appendChild(taskEditIcon);
    taskSettingsBlock.appendChild(taskDeleteIcon);

    taskText.addEventListener('click', ev => {
        taskCheckBox.click();
    });

    taskArrowUp.addEventListener('click', e => {
        let currentTask = taskBlock;
        let order = Number(currentTask.getAttribute("order"));
        let prevTask = currentTask.parentNode.querySelector(".task[order=\"" + (order - 1) + "\"]");
        if (prevTask != null) {
            let xhr = new XMLHttpRequest();
            xhr.open("POST", "/task/priority/up", false);
            xhr.setRequestHeader("Content-Type", "application/json");
            xhr.setRequestHeader("Authorization", "Bearer " + localStorage.getItem("token"));

            let task = {};
            task.id = id;
            task.id_project = idProject;

            xhr.send(JSON.stringify(task));

            let responseObj = JSON.parse(xhr.responseText);

            if (!responseObj.status) {
                alert(responseObj.message);
                return
            }

            currentTask.parentNode.insertBefore(currentTask, prevTask);
            currentTask.setAttribute("order", order - 1);
            prevTask.setAttribute("order", order);
        }

    });

    taskArrowDown.addEventListener('click', e => {

        let currentTask = taskBlock;
        let order = Number(currentTask.getAttribute("order"));
        let nextTask = currentTask.parentNode.querySelector(".task[order=\"" + (order + 1) + "\"]");
        if (nextTask != null) {
            let xhr = new XMLHttpRequest();
            xhr.open("POST", "/task/priority/down", false);
            xhr.setRequestHeader("Content-Type", "application/json");
            xhr.setRequestHeader("Authorization", "Bearer " + localStorage.getItem("token"));
            let task = {};
            task.id = id;
            task.id_project = idProject;
            xhr.send(JSON.stringify(task));

            let responseObj = JSON.parse(xhr.responseText);

            if (!responseObj.status) {
                alert(responseObj.message);
                return
            }

            currentTask.parentNode.insertBefore(nextTask, currentTask);
            currentTask.setAttribute("order", order + 1);
            nextTask.setAttribute("order", order);
        }
    });

    taskDeleteIcon.addEventListener('click', e => {
        if (!confirm("You sure to want delete this task?")) return;

        let order = Number(taskBlock.getAttribute("order"));
        let tasks = taskBlock.parentElement.querySelectorAll(".task");

        let xhr = new XMLHttpRequest();
        xhr.open("DELETE", "/task/delete", false);
        xhr.setRequestHeader("Content-Type", "application/json");
        xhr.setRequestHeader("Authorization", "Bearer " + localStorage.getItem("token"));

        let task = {};
        task.id = id;
        task.id_project = idProject;

        xhr.send(JSON.stringify(task));

        let responseObj = JSON.parse(xhr.responseText);

        if (!responseObj.status) {
            alert(responseObj.message);
            return
        }

        taskBlock.remove();
        for (let i = 0; i < tasks.length; i++) {
            let temp_order = Number(tasks[i].getAttribute("order"));
            if (temp_order > order) {
                tasks[i].setAttribute("order", temp_order - 1)
            }
        }
    });

    taskEditIcon.addEventListener("click", () => {
        if (taskText.style.display !== "none") {
            taskText.style.display = "none";
            taskEditBlock.style.display = "flex";
        }
    });

    taskCheckBox.addEventListener("change", (event) => {
        let temp_checked = taskCheckBox.checked;
        let xhr = new XMLHttpRequest();
        xhr.open("PUT", "/task/check", false);
        xhr.setRequestHeader("Authorization", "Bearer " + localStorage.getItem("token"));
        xhr.setRequestHeader("Content-Type", "application/json");

        let task = {};
        task.id = id;
        task.id_project = idProject;
        task.checked = temp_checked;

        xhr.send(JSON.stringify(task));

        let responseObj = JSON.parse(xhr.responseText);

        if (!responseObj.status) {
            alert(responseObj.message);
            return
        }

        checked = temp_checked;
        taskCheckBox.checked = checked;

    });

    return taskBlock
}


function validateTask(text, deadline) {
    if (text.length === 0 || deadline.length === 0) {
        return false;
    }

    return true;
}


function renderProjectAddBtn() {
    let footer = document.querySelector("footer");
    let html =
        `<div class="create-project" id="create-project-block">
            <i class="far fa-calendar-alt project__header-calendar" id="set-project-deadline-icon"></i>
            <input type="text" placeholder="Start typing here to set title project..." id="project-title-input">
            <button id="create-project-btn">Create Project</button>
        </div>
        <button class="footer__add-project-btn" id="add-project-btn"><i class="fas fa-plus"></i> Add TODO List</button>`;
    footer.innerHTML = html + footer.innerHTML;

    let createProjectBlock = document.getElementById("create-project-block");
    let addProjectBtn = document.getElementById("add-project-btn");
    let createProjectBtn = document.getElementById("create-project-btn");
    let projectTitleInput = document.getElementById("project-title-input");


    addProjectBtn.addEventListener('click', function () {
        createProjectBlock.style.display = "flex";
        projectTitleInput.focus();
    });

    createProjectBtn.addEventListener('click', function () {
        let title = projectTitleInput.value;

        // Send request for add project
        let xhr = new XMLHttpRequest();

        xhr.open('POST', '/project/create', false);
        xhr.setRequestHeader('Content-Type', 'application/json');
        xhr.setRequestHeader("Authorization", "Bearer " + localStorage.getItem("token"));

        let project = {};
        project.name = title;

        xhr.send(JSON.stringify(project));


        let responseObj = JSON.parse(xhr.responseText);


        if (!responseObj.status) {
            alert(responseObj.message);
            return
        }

        renderProject(responseObj.project.id, responseObj.project.name, []);

        projectTitleInput.value = "";
        createProjectBlock.style.display = "none";

        // -----------------------------------------------------------------------------------
    });

    projectTitleInput.addEventListener('keypress', e => {
        if (e.key === "Enter") {
            createProjectBtn.click();
        }
    });
}





