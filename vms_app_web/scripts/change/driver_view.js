const $changeButton = document.querySelector(".c");
const $appView = document.querySelector("#app");
let modal = "";

async function viewFunction(record) {
    elementsView = record.parentNode.parentNode;
    elementView = elementsView.querySelectorAll(".text");

    const user_id = elementView[0].getElementsByTagName("p")[0].innerHTML;

    const response = await fetch("/drivers", {
        headers: {
            "content-type": "application/json",
            accept: "application/json",
        },
    });

    responseBody = await response.json();

    console.log(responseBody);

    responseBody.forEach((onin) => {
        if (onin._id == user_id) {
            updateModal(onin);
        }
    });

    const cancel = document.querySelector("#cancel");

    cancel.onclick = function () {
        modalAdd = document.querySelector(".new-bg").remove();
    };

    const save = document.querySelector("#save");

    //     save.onclick = async function () {
    //         const response = await fetch("http://localhost:8080/users/edit", {
    //             method: "PUT",
    //             body: JSON.stringify({
    //                 user_id: entity[0].user_id,
    //                 given_name: entity[0].given_name,
    //                 email: entity[0].email,
    //                 surname: entity[0].surname,
    //                 password: entity[0].password,
    //                 profile_description: entity[0].profile_description,
    //                 phone_number: entity[0].phone_number,
    //                 city: entity[0].city,
    //             }),
    //             headers: {
    //                 "Content-type": "application/json; charset=UTF-8",
    //                 Accept: "application/json",
    //             },
    //         });
    //         const message = await response.json();
    //         console.log(message);
    //         modalAdd = document.querySelector(".new-bg").remove();
    //     };
}

function updateModal(result) {
    const viewUser = `<div class="new-bg">
        <div class="new-modal">
            <form>
                <div class="input-group">
                    <div class = "inp-row">
                        <div class="input-field">
                            <p>Given Name</p>
                            <input id="given_name" value = "${result.first_name}" onchange="updatedVals(this)" type="text" required />
                        </div>
                        <div class="input-field">
                            <p>Surname</p>
                            <input id="surname" value = "${result.middle_name}" onchange="updatedVals(this)" type="text" required  />
                        </div>
                    </div>
                    <div class = "inp-row">
                        <div class="input-field">
                            <p>Last Name</p>
                            <input id="last_name" value = "${result.last_name}" onchange="updatedVals(this)" type="text" required />
                        </div>
                        <div class="input-field">
                            <p>Goverment ID</p>
                            <input id="goverment_id" value = "${result.goverment_id}" onchange="updatedVals(this)" type="text" required  readonly/>
                        </div>
                    </div>
                    <div class = "inp-row">
                        <div class="input-field">
                            <p>Email</p>
                            <input id="email" value = "${result.email}" onchange="updatedVals(this)"  type="email" required readonly/>
                        </div>
                        <div class="input-field">
                            <p>Phone Number</p>
                            <input id="phone_number"value = "${result.phone_number}" onchange="updatedVals(this)" type="tel" required />
                        </div>
                    </div>

                    <div class = "inp-row under">
                        <div class="input-field">
                            <p>Address</p>
                            <input id="address" value = "${result.address}" onchange="updatedVals(this)" type="text" required />
                        </div>

                        <div class="input-field">
                            <p>Password</p>
                            <input id="password" value = "${result.password}" onchange="updatedVals(this)" type="text" required/>
                        </div>
                    </div>
                    
                    <div class="users-cancelsub">
                        <button type="button" id="cancel" class="addnewbut">
                            <p>Cancel</p>
                        </button>
                        <button type="button"  id="save" class="addnewbut">
                            <p>Edit</p>
                        </button>
                        </button>
                    </div>
                </div>
            </form>
        </div>
    </div>`;
    $appView.insertAdjacentHTML("afterbegin", viewUser);
}
