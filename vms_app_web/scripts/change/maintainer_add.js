const newUser = `<div class="new-bg">
    <div class="new-modal">
        <form>
            <div class="input-group">
                <div class = "inp-row">
                    <div class="input-field">
                        <p>First Name</p>
                        <input id="first_name" type="text" required />
                    </div>
                    <div class="input-field">
                        <p>Last Name</p>
                        <input id="last_name" type="text" required />
                    </div>
                </div>
                
                <div class = "inp-row">
                    <div class="input-field">
                        <p>Email</p>
                        <input id="email" type="email" required />
                    </div>
                    <div class="input-field">
                        <p>Middle Name</p>
                        <input id="middle_name" type="text" required />
                    </div>
                </div>

                <div class = "inp-row under">
                <div class="input-field">
                        <p>Phone Number</p>
                        <input id="phone_number" type="tel" required />
                    </div>
                    
                
                    <div class="input-field">
                        <p>Government Id</p>
                        <input id="government_id" type="number" required />
                    </div>
                </div>

            
                <div class = "inp-row under">
                
                    <div class="input-field">
                        <p>Password</p>
                        <input id="password" type="text" required />
                    </div>
                </div>
                <div class="users-cancelsub">
                    <button type="submit" id="cancel" class="addnewbut">
                        <p>Cancel</p> 
                    </button>
                    <button type="button"  id="save" class="addnewbut">
                        <p>Save</p>
                    </button>
                </div>
            </div>
        </form>
    </div>
</div>`;

const $addButton = document.querySelector(".add-button");
const $app = document.querySelector("#app");

$addButton.onclick = async function () {
    $app.insertAdjacentHTML("afterbegin", newUser);

    const first_name = document.querySelector("#first_name");
    const middle_name = document.querySelector("#middle_name");
    const email = document.querySelector("#email");
    const password = document.querySelector("#password");
    const surname = document.querySelector("#last_name");
    const government_id = document.querySelector("#government_id");
    const phone_number = document.querySelector("#phone_number");
    let modalAdd = "";
    let maintainerID = "";
    const cancel = document.querySelector("#cancel");

    cancel.onclick = function () {
        modalAdd = document.querySelector(".new-bg").remove();
    };

    const save = document.querySelector("#save");
    function isEmpty(value) {
        return value.trim() === "";
    }
    const role_response = await fetch(`/roles`, {
        headers: {
            "Content-type": "application/json; charset=UTF-8",
            Accept: "application/json",
        },
    });
    const resp_roles = await role_response.json();
    console.log(resp_roles);
    resp_roles.forEach((oneinstance) => {
        if (oneinstance.name == "maintenance person") {
            maintainerID = oneinstance._id;
            console.log(maintainerID);
        }
    });

    save.onclick = async function () {
        if (
            isEmpty(first_name.value) ||
            isEmpty(email.value) ||
            isEmpty(password.value) ||
            isEmpty(surname.value) ||
            isEmpty(phone_number.value) ||
            isEmpty(government_id.value)
        ) {
            text = "ivalid input";
            alert(text);
        } else {
            const response = await fetch("/createUser", {
                method: "POST",
                body: JSON.stringify({
                    role_id: maintainerID,
                    first_name: first_name.value,
                    middle_name: surname.value,
                    email: email.value,
                    last_name: surname.value,
                    password: password.value,
                    goverment_id: parseInt(government_id.value),
                    phone_number: phone_number.value,
                }),
                headers: {
                    "Content-type": "application/json; charset=UTF-8",
                    Accept: "application/json",
                },
            });
            const message = await response.json();
            console.log(message);

            modalAdd = document.querySelector(".new-bg").remove();
        }
        return false;
    };
};
