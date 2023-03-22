import React, {useEffect, useState} from 'react'

import '../css/MedicationRequestForm.css'
import { resolveURL } from "../util/resolveURL";

const MedicationRequestForm = () => {

    const [medications, setMedications] = useState([])
    const [medicationListVisibility, showMedicationList] = useState(false)

    const [drugList, setDrugList] = useState([])
    const [drugListVisibility, showDrugList] = useState(false)
    const [drugListEmpty, showDrugListEmpty] = useState(false)

    const [submitVisbility, showSubmitButton] = useState(false)
    const [submitText, setSubmitText] = useState("Submit Medical Record")

    const [consentID, setConsentID] = useState("")
    const [patientName, setPatientName] = useState("")
    const [drugName, setDrugName] = useState("")
    const [medicationNote, setMedicationNote] = useState("")
    const [effectiveDateTime, setEffectiveDateTime] = useState("")
    const [dosageSequence, setDosageSequence] = useState(0)
    const [dosageNote, setDosageNote] = useState("")
    const [medicationStatus, setMedicationStatus] = useState("active")

    const [errorMessage, setErrorMessage] = useState("")
    const [errorMessageVisibility, showErrorMessage] = useState(false)

    useEffect(async () => {
        let param = window.location.search.split('consent_id=')[1]
        setConsentID(param)

        let cookieData = await getCookie();
        console.log(cookieData);

        if (cookieData.message !== "success") {
            window.location = "/not-authorized";
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

    // Clear form, add medication to medications state.
    const addMedication = async () => {
        if (consentID !== ""
            && patientName !== ""
            && drugName !== ""
            && medicationNote !== ""
            && effectiveDateTime !== ""
            && dosageSequence !== null
            && dosageNote !== ""
            && medicationStatus !== "" ) {

            let validateDrugName = await drugExists(drugName)
            console.log(validateDrugName)
            if (validateDrugName.message === 'valid' ) {
                let medication = {
                    consentid: consentID,
                    name: patientName,
                    medication: drugName,
                    note: medicationNote,
                    effectivedatetime: effectiveDateTime,
                    dosagesequence: dosageSequence,
                    dosagenote: dosageNote,
                    status: medicationStatus
                }

                console.log(medication)
                
                medications.push(medication)
                clearForm()
    
                showMedicationList(true)
                showSubmitButton(true)
            } else {
                setErrorMessage("Please add a valid drug name!")
                showErrorMessage(true)
    
                setTimeout(() => {
                    showErrorMessage(false)
                }, 3000)
            }

        } else {
            setErrorMessage("Please fill all fields!")
            showErrorMessage(true)

            setTimeout(() => {
                showErrorMessage(false)
            }, 3000)
        }
    }

    const clearForm = () => {
        setDrugName("")
        setMedicationNote("")
        setEffectiveDateTime("")
        setDosageSequence(0)
        setDosageNote("")
        setMedicationStatus("active")
    }

    // Delete medication from current history
    const deleteMedication = (med_elem) => {
        let filtered = medications.filter(elem => JSON.stringify(elem) !== JSON.stringify(med_elem) )

        setMedications(filtered)
        if (filtered.length === 0) {
            showMedicationList(false)
            showSubmitButton(false)
        } else {
            showMedicationList(true)
        }
    }

    const handleSearch = async (e) => {
        let input = e.target.value
        setDrugName(input)

        if (input.length > 2) {
            let data = await getDrugData(input)
            setDrugList(data)
            showDrugList(true)

            data.length === 0 ? showDrugListEmpty(true) : showDrugListEmpty(false)
        } else {
            setDrugList([])
            showDrugList(false)
        }
    }

    const selectDrug = (drugName) => {
        setDrugName(drugName)
        setDrugList([])
        showDrugList(false)
    }

    const getDrugData = async (input) => {
        let url = `${resolveURL("standardization")}/standardization/${input}`;
        const response = await fetch(url, {
            method: "GET",
            mode: "cors",
            credentials: "include"
        })
    
        return response.json()
    }

    const drugExists = async (input) => {
        let url = `${resolveURL("standardization")}/standardization/valid/${input}`;
        const response = await fetch(url, {
            method: "GET",
            mode: "cors",
            credentials: "include"
        })
    
        return response.json()
    }

    const postMedicalForm = async (data) => {
        let url = `${resolveURL("medication-record")}/medication-record/new`
        const response = await fetch(url, {
            method: "POST",
            mode: "cors",
            credentials: "include",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
        })

        return response.json()
    }

    // Fetch POST medication history to backend
    const submitMedicationHistory = async () => {
        console.log(medications)

        let final_medications = {
            records: medications
        }

        let response = await postMedicalForm(final_medications)
        
        setSubmitText("Medication Added!")
        setTimeout(function() {
            window.location.href = '/inbox'
        }, 2000)


    }

    return(<div>
        <div className='medication-request-form'>
            <p id='title'>Medication Form</p>
                { errorMessageVisibility && (
                    <p id='form-error'>{errorMessage}</p>
                ) }
                <label htmlFor="constentID">
                <span>Enter the Consent Id</span>
                    <input disabled type="text" name="constentID" placeholder='Consent Id' value={consentID} onChange={e => setConsentID(e.target.value)}/>
                </label>
                <label htmlFor="patientName">
                <span>Patient's Full Name</span>
                    <input type="text" name="patientName" placeholder='Patient Full Name' value={patientName} onChange={e => setPatientName(e.target.value)}/>
                </label>
                <label htmlFor="drugName">
                <span>Drug Name</span>
                    <input type="text" name="drugName" placeholder='Drug name' autoComplete='off' value={drugName} onChange={e => handleSearch(e)}/>
                </label>

                {drugListEmpty && (
                    (<p id='drug-list-empty'>No drugs found</p>)
                )}

                {drugListVisibility && drugList && (
                    <div className='drug-list'>
                        {drugList.map((drug, idx) => {
                            return (<li onClick={() => selectDrug(drug)} key={idx}>{drug}</li>)
                        })}
                    </div>
                )}

                <label htmlFor="medicationNote">
                <span>Medication Note</span>
                    <textarea name='medicationNote' placeholder='Medication Note' value={medicationNote} onChange={e => setMedicationNote(e.target.value)}></textarea>
                </label>

                <label htmlFor="effectiveDateTime">
                <span>Effective Date Time</span>
                    <input type="date" name='effectiveDateTime' placeholder='Date of when pill was first prescribed (i.e. yyyy-mm-dd)' value={effectiveDateTime} onChange={e => setEffectiveDateTime(e.target.value)}></input>
                </label>

                <label htmlFor="dosageSequence">
                <span>Dosage Sequence</span>
                    <input type="number" name='dosageSequence' placeholder='Sequence of how many times patient takes medication (i.e. 1)' value={dosageSequence} onChange={e => setDosageSequence(e.target.value)}></input>
                </label>

                <label htmlFor="dosageNote">
                <span>Dosage Note</span>
                    <textarea name='dosageNote' placeholder='Note specifying dosage' value={dosageNote} onChange={e => setDosageNote(e.target.value)}></textarea>
                </label>

                <label htmlFor="medicationStatus">
                <span>Medication Status</span>
                    <select name="medicationStatus" id="medicationStatus" value={medicationStatus} onChange={e => setMedicationStatus(e.target.value)}>
                        <option value="active">Active</option>
                        <option value="inactive">Stopped</option>
                        <option value="finished">Finished</option>
                    </select>
                </label>

                <button id='add-medication' onClick={addMedication}>Add Medication</button>
        </div>
                    
        {medicationListVisibility && medications && (
            <div className='added-medications'>
                <p id='title'>Added Medications</p>
                {medications.map((med_elem, idx) => {
                    return (<div className='medication-elem' id={med_elem.medication + "-" + idx} key={idx}>
                        <p id='information'>{med_elem.effectivedatetime} | {med_elem.medication}</p>
                        <button id='delete' onClick={() => {deleteMedication(med_elem)}}>Delete</button>
                    </div>)
                })}
            </div>
        )}

        {submitVisbility && (
            <button id='submit-medical-record' onClick={() => {submitMedicationHistory()}}>{submitText}</button>
        )}
    </div>)

}

export default MedicationRequestForm