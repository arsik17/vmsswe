const getUrl = window.location;
const baseUrl =
    getUrl.protocol + "//" + getUrl.host + "/" + getUrl.pathname.split("/")[1];

const signInBtn = document.querySelector("#sign-in");
const email = document.querySelector("#email");
const password = document.querySelector("#password");

const forgotBtn = document.querySelector("#forgot");
const showBtn = document.querySelector("#show");

showBtn.onclick = signInBtn.onclick = function register() {
    fetch("/loginAdmin", {
        method: "post",
        body: JSON.stringify(formData),
        mode: "cors",
    }).then((response) => {
        if (response.ok) {
            alert("You have logged in successfuly");
            //window.location.href = "./login"
        } else {
            throw "unauthorised";
        }
    }); //TODO: catch throw
    return false;
};
