import {
	collectFormValues,
	collectReviewValues,
	validateNairobiTourFormV,
	displayFormErrors,
	displayReviewFormErrors,
	validateCarHireForm,
	validateEnquiryForm,
	validateReviewForm,
} from "./taxi-form.js";

function initContact(form) {
	if (form === "nairobi") {
		document
			.getElementById("contact-form-submit")
			.addEventListener("click", () => {
				const entries = collectFormValues();
				const errors = validateNairobiTourFormV(entries);
				displayFormErrors(errors);

				// send post request to my backend
				if (errors.length === 0) {
					sendContactForm(entries, "/contact-form/nairobi-tour");
					return;
				}
			});
	} else if (form === "safari") {
		document
			.getElementById("contact-form-submit")
			.addEventListener("click", () => {
				const entries = collectFormValues();
				const errors = validateNairobiTourFormV(entries);
				displayFormErrors(errors);

				// send post request to my backend
				if (errors.length === 0) {
					sendContactForm(entries, "/contact-form/safari");
					return;
				}
			});
	} else if (form === "hire") {
		document
			.getElementById("contact-form-submit")
			.addEventListener("click", () => {
				const entries = collectFormValues();
				const errors = validateCarHireForm(entries);
				displayFormErrors(errors);

				// send post request to my backend
				if (errors.length === 0) {
					sendContactForm(entries, "/contact-form/car-hire");
					return;
				}
			});
	} else {
		document
			.getElementById("contact-form-submit")
			.addEventListener("click", () => {
				const entries = collectFormValues();
				const errors = validateEnquiryForm(entries);
				displayFormErrors(errors);

				// send post request to my backend
				if (errors.length === 0) {
					sendContactForm(entries, "/contact-form/enquery");
					return;
				}
			});

		document.getElementById("review-btn").addEventListener("click", () => {
			const entries = collectReviewValues();
			const errors = validateReviewForm(entries);
			displayReviewFormErrors(errors);

			// send post request to my backend
			if (errors.length === 0) {
				sendContactForm(entries, "/contact-form/review");
				return;
			}
		});
	}
}

initContact(document.querySelector(".js-form-title").dataset.formName);

function sendContactForm(entries, url) {
	if (entries.serviceEnguery) {
		switch (entries.serviceEnguery) {
			case "0":
				entries.serviceEnguery = "safari";
				break;
			case "1":
				entries.serviceEnguery = "get-a-taxi";
				break;
			case "2":
				entries.serviceEnguery = "nairobi-tour";
				break;
			case "3":
				entries.serviceEnguery = "car-hire";
				break;
		}
	}

	console.log(entries + " : " + JSON.stringify(entries));

	const body = {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify(entries),
	};
	fetch(url, body)
		.then((response) => response.json())
		.then((data) => {
			console.log("parsed data is: " + data);
			window.location.href = "/";
		})
		.catch((error) => {
			console.error("There was an error!", error);
		});
}
