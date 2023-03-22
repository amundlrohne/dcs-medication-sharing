import React, {useEffect, useState} from 'react'

import '../css/ViewMedicationRecord.css'

import { resolveURL } from "../util/resolveURL";

const ViewMedicationRecord = () => {

    const [consentID, setConsentID] = useState("")
    const [patientName, setPatientName] = useState("")
    const [records, setRecords] = useState([])

    useEffect(async () => {
        let cookieData = await getCookie();
        console.log(cookieData);

        if (cookieData.message !== "success") {
            window.location = "/not-authorized";
        }

        let consent_id = window.location.search.split('consent_id=')[1]

        if (consent_id !== undefined) {
            console.log(consent_id)
            setConsentID(consent_id)
    
            let consentObj = {
                consentid: consent_id
            }
    
            let bundle = await getBundle(consentObj)
            let parsedJSON = JSON.parse(bundle.data.data)
    
            setRecords(parsedJSON.entry[0].resource.entry)
            setPatientName(parsedJSON.entry[0].resource.entry[0].resource.subject.display)
        }
    }, []);

    const getCookie = async () => {
        let url = `${resolveURL("healthcare-provider")}/healthcare-provider/current`;
        const response = await fetch(url, {
            method: "GET",
            mode: "cors",
            credentials: "include",
        });
    
        return response.json();
    };

    const getBundle = async (data) => {
        let url = `${resolveURL("medication-record")}/medication-record`
        const response = await fetch(url, {
            method: "POST",
            mode: "cors",
            credentials: "include",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
        });

        return response.json()
    }

    return (<div>
        <div className='medication-record-container'>
            <div className='primary-information'>
                <label htmlFor="constentID">
                <span>Consent ID</span>
                    <input disabled type="text" name="constentID" value={consentID}/>
                </label>
                <label htmlFor="patientName">
                <span>Patient Name</span>
                    <input disabled type="text" name="patientName" value={patientName}/>
                </label>
            </div>

            {records.map((record_elem, idx) => {
                            return (<div className='record-elem' id={consentID} key={idx}>
                                        <p>Medication #{idx+1}</p>
                                        <label htmlFor="medication">
                                        <span>Medication</span>
                                            <input disabled type="text" name="medication" value={record_elem.resource.medicationCodeableConcept.text}/>
                                        </label>

                                        <label htmlFor="effectivedatetime">
                                        <span>Effective Date Time</span>
                                            <input disabled type="text" name="effectivedatetime" value={record_elem.resource.effectiveDateTime}/>
                                        </label>

                                        <label htmlFor="doctorsnote">
                                        <span>Medication Note</span>
                                            <textarea disabled type="text" name="doctorsnote" value={record_elem.resource.note[0].text}/>
                                        </label>

                                        <label htmlFor="status">
                                        <span>Status</span>
                                            <input disabled type="text" name="status" value={record_elem.resource.status}/>
                                        </label>

                                        <label htmlFor="dosagesequence">
                                        <span>Dosage Sequence</span>
                                            <input disabled type="number" name="dosagesequence" value={record_elem.resource.dosage[0].sequence}/>
                                        </label>

                                        <label htmlFor="dosagenote">
                                        <span>Dosage Note</span>
                                            <textarea disabled type="text" name="dosagenote" value={record_elem.resource.dosage[0].text}/>
                                        </label>
                                    </div>)
                            })}
        </div>
    </div>)
}

export default ViewMedicationRecord