function renderLoginForm() {
    document.querySelector("#projects-list").innerHTML = `
    <div class="auth-block"> 
        <form name="login_form">
        <h4>Login</h4>
            <input type="email" name="email" required placeholder="Email">
            <input type="password" name="password" required placeholder="Password">
            <div>    
                <button type="submit" class="login-btn">Login</button>
                <button type="button" id="registration-form-btn">Registration</button>
            </div>
        </form>
    </div>
    `;

    let form = document.forms.login_form;

    form.addEventListener("submit", (event) => {
        event.preventDefault();

        let email = form.email.value;
        let password = form.password.value;

        if (!email.includes("@")) {
            alert("Email address is not valid");
            form.email.focus();
            return
        }

        if (password.length < 6) {
            alert("Password must be a least 6 characters");
            form.password.focus();
            return
        }


        let requestObj = {};
        requestObj.email = email;
        requestObj.password = password;

        let xhr = new XMLHttpRequest();

        xhr.open("POST", "/user/login", false);
        xhr.setRequestHeader("Content-Type", "application/json");
        xhr.send(JSON.stringify(requestObj));

        let responseObj = JSON.parse(xhr.responseText);

        if (!responseObj.status) {
            alert(responseObj.message);
            return
        }

        localStorage.setItem("token", responseObj.account.token);

        document.querySelector(".auth-block").remove();

        getProjects();
        renderProjectAddBtn();
    });


    let registrationBtn = document.querySelector("#registration-form-btn");

    registrationBtn.addEventListener("click", () => {
        renderRegistrationForm();
    });


}


function renderRegistrationForm() {

    document.querySelector("#projects-list").innerHTML = `
    <div class="auth-block"> 
        <form name="registration_form">
        <h4>Registration</h4>
            <input type="email" name="email" required placeholder="Email">
            <input type="password" name="password" required placeholder="Password">
            <input type="password" name="password2" required placeholder="Confirm password">
            <div>    
                <button type="button" id="login-form-btn">Login</button>
                <button type="submit" id="registration-form-btn">Registration</button>
            </div>
        </form>
    </div>
    `;

    let reg_form = document.forms.registration_form;

    reg_form.addEventListener("submit", (event) => {
        event.preventDefault();

        let email = reg_form.email.value;
        let password = reg_form.password.value;
        let password2 = reg_form.password2.value;

        if (!email.includes("@")) {
            alert("Email address is not valid");
            reg_form.email.focus();
            return
        }

        if (password.length < 6) {
            alert("Password must be a least 6 characters");
            reg_form.password.focus();
            return
        }

        if (password !== password2) {
            alert("Passwords does not matched!");
            reg_form.password2.focus();
            return
        }

        let requestObj = {};
        requestObj.email = email;
        requestObj.password = password;

        let xhr = new XMLHttpRequest();

        xhr.open("POST", "/user/registration", false);
        xhr.setRequestHeader("Content-Type", "application/json");
        xhr.send(JSON.stringify(requestObj));

        let responseObj = JSON.parse(xhr.responseText);

        console.log(responseObj);

        if (!responseObj.status) {
            alert(responseObj.message);
            return
        }

        localStorage.setItem("token", responseObj.account.token);

        document.querySelector(".auth-block").remove();

        alert("You have successfully registered. Remember your email and password");

        renderProjectAddBtn();
    });

    let loginBtn = document.querySelector("#login-form-btn");

    loginBtn.addEventListener("click", () =>{
       renderLoginForm();
    });
}


