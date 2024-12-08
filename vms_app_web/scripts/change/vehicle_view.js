const $changeButton = document.querySelector(".c");
const $appView = document.querySelector("#app");
let modal = "";

async function viewFunction(record) {
    elementsView = record.parentNode.parentNode;
    elementView = elementsView.querySelectorAll(".text");

    const user_id = elementView[0].getElementsByTagName("p")[0].innerHTML;

    const response = await fetch("/vehicles", {
        headers: {
            "content-type": "application/json",
            accept: "application/json",
        },
    });

    responseBody = await response.json();

    console.log(responseBody);

    responseBody.forEach((onin) => {
        if (onin.vin == user_id) {
            updateModal(onin);
        }
    });

    const cancel = document.querySelector("#cancel");

    cancel.onclick = function () {
        modalAdd = document.querySelector(".new-bg").remove();
    };

    //     const save = document.querySelector("#save");

    //     //     save.onclick = async function () {
    //     //         const response = await fetch("http://localhost:8080/users/edit", {
    //     //             method: "PUT",
    //     //             body: JSON.stringify({
    //     //                 user_id: entity[0].user_id,
    //     //                 given_name: entity[0].given_name,
    //     //                 email: entity[0].email,
    //     //                 surname: entity[0].surname,
    //     //                 password: entity[0].password,
    //     //                 profile_description: entity[0].profile_description,
    //     //                 phone_number: entity[0].phone_number,
    //     //                 city: entity[0].city,
    //     //             }),
    //     //             headers: {
    //     //                 "Content-type": "application/json; charset=UTF-8",
    //     //                 Accept: "application/json",
    //     //             },
    //     //         });
    //     //         const message = await response.json();
    //     //         console.log(message);
    //     //         modalAdd = document.querySelector(".new-bg").remove();
    //     //     };
}

function updateModal(result) {
    const viewUser = `<div class="new-bg">
        <div class="new-modal1">
            <form>
                <div class="input-group">
                    <div class = "inp-row">
                        <div class="input-field">
                            <p>Driver ID</p>
                            <input id="driver_id" value = "${result.driver_id}" onchange="updatedVals(this)" type="text" required />
                        </div>
                        <div class="input-field">
                            <p>Car Make</p>
                            <input id="car_make" value = "${result.car_make}" onchange="updatedVals(this)" type="text" required  />
                        </div>
                    </div>
                    <div class = "inp-row">
                        <div class="input-field">
                            <p>Model</p>
                            <input id="model" value = "${result.model}" onchange="updatedVals(this)" type="text" required />
                        </div>
                        <div class="input-field">
                            <p>License Plate</p>
                            <input id="license_plate" value = "${result.license_plate}" onchange="updatedVals(this)" type="text" required/>
                        </div>
                    </div>
                    <div class = "inp-row">
                        <div class="input-field">
                            <p>Production Year</p>
                            <input id="production_year" value = "${result.production_year}" onchange="updatedVals(this)"  type="email" required/>
                        </div>
                        <div class="input-field">
                        <p>Status</p>
                        <select name="status" id="inp" required>
                            <option>active</option>
                            <option>inactive</option>
                            <option>other</option>
                            </select>
                            
                        </div>
                    </div>

                    <div class = "inp-row under">
                        <div class="input-field">
                            <p>Exact Location</p>
                            <input id="exact_location" value = "${result.exact_location}" onchange="updatedVals(this)" type="text" required readonly/>
                        </div>

                        <div class="input-field">
                            <p>Mileage</p>
                            <input id="mileage" value = "${result.mileage}" onchange="updatedVals(this)" type="text" required/>
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
