function addModal() {
    const newUser = `<div class="new-bg">
        <div class="new-modal">
            <form>
                <div class="input-group1">
                    <div class = "inp-col left">
                        <div class="input-field">
                            <p>Driver ID</p>
                            <input id="driver_id" type="text" required/>
                            <p id="result">Driver: John Doe<\p>
                        </div>
                        <div class="input-field">
                            <p>Year</p>
                            <input id="year" type="text" required/>
                        </div>
                        <div class="input-field">
                            <p>License plate number</p>
                            <input id="lpn" type="text" />
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
                   
                    <div class = "inp-col">
                        <div class="input-field">
                            <p>VIN</p>
                            <input id="vin" type="text" required/>
                        </div>
                        <div class="input-field">
                            <p>Make</p>
                            <input id="make" type="text" required/>
                        </div>
                        <div class="input-field">
                            <p>Model</p>
                            <input id="model" type="text" required/>
                        </div>
                        
                    
                        <div class="input-field">
                            <p>Mileage in thouthands</p>
                            <input id="mileage" type="number" required/>
                        </div> 

                        
                    </div>

                    
                </div>
                
                    <div class="users-cancelsub">
                        <button type="submit" id="cancel" class="addnewbut">
                            <p>Cancel</p>
                        </button>
                        <button type="button"  id="save" class="addnewbut">
                            <p>Save</p>
                        </button>
                        </button>
                    </div>
            </form>
        </div>
    </div>`;
    $app.insertAdjacentHTML("afterbegin", newUser);
    // <p id="select-img">Select Image</p>
    //<input type="file" id="avatar" name="avatar" accept="image/png, image/jpeg" />
}

const $addButton = document.querySelector(".add-button");
const $app = document.querySelector("#app");
const $driverResult = document.getElementById("#result");

$addButton.onclick = async function () {
    // let driverName = document.getElementById("result");
    // driverName = "ADADAD";
    addModal();
    userEntity = {};
    newEntity = {};

    // caregiver_user_id.onchange = async function () {
    //     let response_driver = await fetch(`/drivers`, {
    //         headers: {
    //             "Content-type": "application/json; charset=UTF-8",
    //             Accept: "application/json",
    //         },
    //     });
    //     newEntity = await response_driver.json();

    //     const url = "http://localhost:8080/users/one";
    //     response = await fetch(`${url}/${caregiver_user_id.value}`, {
    //         headers: {
    //             "Content-type": "application/json; charset=UTF-8",
    //             Accept: "application/json",
    //         },
    //     });
    //     userEntity = await response.json();

    //     const surname = document.querySelector("#surname");
    //     const given_name = document.querySelector("#given_name");

    //     if (newEntity[0] != undefined) {
    //         if (newEntity[0].caregiver_user_id == userEntity[0].user_id) {
    //             given_name.value = "Caregiver already exists";
    //             surname.value = "Caregiver already exists";
    //             ready = false;
    //         }
    //     } else {
    //         if (userEntity[0] != undefined) {
    //             ready = true;
    //             given_name.value = userEntity[0].given_name;
    //             surname.value = userEntity[0].surname;
    //         } else {
    //             ready = false;
    //             given_name.value = "No such user exists";
    //             console.log(given_name);
    //             surname.value = "No such user exists";
    //         }
    //     }
    // };

    cancel.onclick = function () {
        modalAdd = document.querySelector(".new-bg").remove();
    };

    const year = document.querySelector("#year");
    const driver_id = document.querySelector("#driver_id");
    const lpn = document.querySelector("#lpn");
    const status = document.getElementById("inp");
    const vin = document.querySelector("#vin");
    const make = document.querySelector("#make");
    const model = document.querySelector("#model");
    const mileage = document.querySelector("#mileage");

    function isEmpty(value) {
        return value.trim() === "";
    }

    save.onclick = async function () {
        if (
            isEmpty(year.value) ||
            isEmpty(driver_id.value) ||
            isEmpty(lpn.value) ||
            isEmpty(status.value) ||
            isEmpty(vin.value) ||
            isEmpty(model.value) ||
            isEmpty(mileage.value) ||
            isEmpty(make.value)
        ) {
            text = "ivalid input";
            alert(text);
        } else {
            const response1 = await fetch("/createVehicle", {
                method: "POST",
                body: JSON.stringify({
                    _id: vin.value,
                    production_year: parseInt(year.value),
                    mileage: parseInt(mileage.value),
                    model: model.value,
                    license_plate: lpn.value,
                    driver_id: driver_id.value,
                    car_make: make.value,
                    activity_status: status.value,
                    exact_location: "0.0;0.0",
                }),
                headers: {
                    "Content-type": "application/json; charset=UTF-8",
                    Accept: "application/json",
                },
            });
            const message1 = await response1.json();

            modalAdd = document.querySelector(".new-bg").remove();
        }
    };
};
