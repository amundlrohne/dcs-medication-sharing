import React from "react";

import "../css/MedicationForm.css";
import { resolveURL } from "../util/resolveURL";

class MedicationForm extends React.Component {
    constructor() {
        super();
        this.states = {
            textareaHeight: "",
        };
    }

    // Search drug name from standardization service
    searchTest(event) {
        if (
            (event.keyCode >= 48 && event.keyCode <= 57) ||
            (event.keyCode >= 65 && event.keyCode <= 90)
        ) {
            let s = event.target.value + event.key;
            //console.log(s);

            if (s > 2) {
                //Testing
                let data = ["data1", "dawdawd", "dawdawd"];

                // Change url
                fetch(`${resolveURL("standardization")}/standardization/${s}`)
                    .then((response) => response.json())
                    .then(
                        (data) => console.log(data)
                        // Propose data as options
                    );
            }
        }
    }

    render() {
        return (
            <div className="MedicationForm">
                <form className="med-form" action="" method="post">
                    <label for="constentID">
                        <span>Enter the Consent Id</span>
                        <input
                            type="text"
                            name="constentID"
                            placeholder="Consent Id"
                            required="true"
                        />
                    </label>
                    <label for="patientName">
                        <span>Patient's Full Name</span>
                        <input
                            type="text"
                            name="patientName"
                            placeholder="Patient Full Name"
                            required="true"
                        />
                    </label>
                    <label for="drugName">
                        <span>Drug Name</span>
                        <input
                            type="text"
                            name="drugName"
                            placeholder="Drug name"
                            required="true"
                            onKeyDown={this.searchTest}
                        />
                    </label>

                    <label for="medicationNote">
                        <span>Medication Note</span>
                        <textarea
                            name="mediationNote"
                            placeholder="Medication Note"
                            required="true"
                        ></textarea>
                    </label>

                    <label for="effectiveDateTime">
                        <span>Effective Date Time</span>
                        <textarea
                            name="effectiveDateTime"
                            placeholder="Date of when pill was first prescribed"
                            required="true"
                        ></textarea>
                    </label>

                    <label for="dosageSequence">
                        <span>Dosage Sequence</span>
                        <textarea
                            name="dosageSequence"
                            placeholder="Sequence of how many times patient takes medication"
                            required="true"
                        ></textarea>
                    </label>

                    <label for="dosageNote">
                        <span>Dosage Note</span>
                        <textarea
                            name="dosageNote"
                            placeholder="Note specifying dosage"
                            required="true"
                        ></textarea>
                    </label>

                    <label for="medicationStatus">
                        <span>Medication Status</span>
                        <select name="medicationStatus" id="medicationStatus">
                            <option value="active">Active</option>
                            <option value="inactive">Inactive</option>
                        </select>
                    </label>

                    <span></span>
                    <input type="submit" value="Send Medication" />
                </form>
            </div>
        );
    }
}

export default MedicationForm;
