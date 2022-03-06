'use strict';

// Divs
let formDiv = document.querySelector("#form-dates")
let singleDateDiv = formDiv.querySelector("#single-date")
let dateRangeDiv = formDiv.querySelector("#date-range")
let multipleWeeksSelectDiv = formDiv.querySelector("#multiple-weeks-select")
let singleWeekSelectDiv = formDiv.querySelector("#single-weeks-select")
let submitBtnDiv = formDiv.querySelector("#submit-container")
let responseDiv = document.querySelector("#response")
// Buttons
let submitBtn = formDiv.querySelector("[submit-btn]")
let multipleWeeksBtn = formDiv.querySelector("[multiple-weeks-btn]")
let singleWeeksBtn = formDiv.querySelector("[single-weeks-btn]")

// Inputs
let singleDate = formDiv.querySelector("input[name='date-in']") 
let rgStartDate = formDiv.querySelector("input[name='start-date']")
let rgEndDate = formDiv.querySelector("input[name='end-date']")

let isDateRange = false;

// Init data to send
let data = {}

multipleWeeksBtn.addEventListener("click", (event) => {
    event.preventDefault()
    dateRangeDiv.classList.remove("hidden")
    submitBtnDiv.classList.remove("hidden")
    singleWeekSelectDiv.classList.remove("hidden")

    singleDateDiv.classList.add("hidden")
    multipleWeeksSelectDiv.classList.add("hidden")
    isDateRange = true;
})

singleWeeksBtn.addEventListener("click", (event) => {
    event.preventDefault()
    dateRangeDiv.classList.add("hidden")
    singleWeekSelectDiv.classList.add("hidden")
    singleDateDiv.classList.remove("hidden")
    multipleWeeksSelectDiv.classList.remove("hidden")
    isDateRange = false;
})

submitBtn.addEventListener("click", (event) => {
    event.preventDefault()
    let startDate = isDateRange ? new Date(rgStartDate.value) : null
    let endDate = isDateRange ? new Date(rgEndDate.value) : null
    let singleDateDate = isDateRange ? null : new Date(singleDate.value)
    data = {
        "dateRange": isDateRange ? {
            "startDate": {
                "year": startDate.getFullYear(),
                "day": startDate.getDate(),
                "month": startDate.getMonth()+1,
            },
            "endDate": {
                "year": endDate.getFullYear(),
                "day": endDate.getDate(),
                "month": startDate.getMonth()+1,
            },
        } : null,
        "singleDate": !isDateRange ? 
            {
                "year": singleDateDate.getFullYear(),
                "day": singleDateDate.getDate(),
                "month": singleDateDate.getMonth()+1,
            } : null
    }
    let res = submitDates("/date", data)
    parseResponse(res) 
    console.log(data)
    /**
    isDateRange && console.log("Date values:", rgStartDate.value, rgEndDate.value)
    !isDateRange && console.log("Date value:", singleDate.value)
    **/
})

async function submitDates(url = '/date', data = {}) {
    const response = await fetch(url, {
        method:'POST',
        credentials: 'same-origin',
        body: JSON.stringify(data)
    });
    var res = await response.json()
    return res;
} 

async function parseResponse(res) {
    let response = await res
    let newNode = !document.querySelector("#error-msg") ? document.createElement("p") : document.querySelector("#error-msg") 
    newNode.setAttribute("id","error-msg")
    if (response.error != null) {
        console.log("error:", response.error)
        responseDiv.classList.remove("success")
        responseDiv.classList.add("failure")
        newNode.innerHTML = "Error pujant dates."
        responseDiv.appendChild(newNode)
        responseDiv.classList.remove("hidden")
        return
    }
    console.log(response)
    responseDiv.classList.remove("failure")
    responseDiv.classList.add("success")
    newNode.innerHTML = "Data pujada satisfactoriament."
    responseDiv.appendChild(newNode)
    responseDiv.classList.remove("hidden")
}