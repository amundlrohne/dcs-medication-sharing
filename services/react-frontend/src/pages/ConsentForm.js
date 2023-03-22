import React, { useEffect, useState } from "react";

import { resolveURL } from "../util/resolveURL";

import "../css/ConsentForm.css";

const ConsentForm = () => {
    const [senderHP, setSenderHP] = useState(["sHP1", "sHP2"]);
    const [receiverHP, setReceiverHP] = useState(["RHP1", "sHP1"]);
    const [selectedSenderHP, setSelectedSenderHP] = useState([]);
    const [selectedReceiverHP, setSelectedReceiverHP] = useState([]);
    const [errorMsg, setErrorMsg] = useState("");
    const [errorMsgVisibility, showErrorMsg] = useState(false);

    const [consentID, setConsentID] = useState("");
    const [consentCreated, showConsentCreated] = useState(false);

    function handleSubmit(e) {
        let senderHP = e.target.senderHP.value;
        let receiverHP = e.target.receiverHP.value;
        if (senderHP === receiverHP) {
            e.preventDefault();
            setErrorMsg(
                "Sender and receiver healthcare providers must be different"
            );
            showErrorMsg(true);

            setTimeout(function () {
                showErrorMsg(false);
            }, 3000);
            return;
        }

        e.preventDefault();
        createConcent();
    }

    // get the id
    async function getHPByName(name) {
        let fetchUrl =
            `${resolveURL("healthcare-provider")}/healthcare-provider/name/` +
            name;
        const res = await fetch(fetchUrl);
        const d = await res.json();
        return d.data.data;
    }

    async function postConsent(senderPK, receiverPK) {
        let postUrl = `${resolveURL("consent")}/consent/`;

        let consentObj = {
            topublickey: receiverPK,
            frompublickey: senderPK,
            expdate: dateNowPlusOneMonth(),
            datecreated: dateNow(),
        };

        console.log(consentObj);

        let finalized_consent_id = await createConsent(postUrl, consentObj);
        setConsentID(finalized_consent_id.data.data.InsertedID);
        showConsentCreated(true);
    }

    const createConsent = async (postUrl, consentObj) => {
        let req = await fetch(postUrl, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(consentObj),
        });

        return req.json();
    };

    function dateNow() {
        let today = new Date();
        let dd = today.getDate();
        let mm = today.getMonth() + 1;
        let yyyy = today.getFullYear();
        if (dd < 10) {
            dd = "0" + dd;
        }

        if (mm < 10) {
            mm = "0" + mm;
        }

        today = yyyy + "-" + mm + "-" + dd;
        return today;
    }

    function dateNowPlusOneMonth() {
        let now = new Date();
        let current = new Date(now.getFullYear(), now.getMonth() + 1, 1);

        let dd = current.getDate();
        let mm = current.getMonth() + 1;
        let yyyy = current.getFullYear();
        if (dd < 10) {
            dd = "0" + dd;
        }

        if (mm < 10) {
            mm = "0" + mm;
        }
        let todayPlusOneMonth = yyyy + "-" + mm + "-" + dd;
        return todayPlusOneMonth;
    }

    useEffect(() => {
        async function fetchHP() {
            //Change url
            let fetchUrl = `${resolveURL(
                "healthcare-provider"
            )}/healthcare-provider/all`;
            let healthCareProviderNames = [];
            let d = [];

            const res = await fetch(fetchUrl);
            d = await res.json();
            d = d.data.data;
            for (let i = 0; i < d.length; i++) {
                healthCareProviderNames.push(d[i].name);
            }
            setSenderHP(healthCareProviderNames);
            setReceiverHP(healthCareProviderNames);
            setSelectedSenderHP(d[0].name);
            setSelectedReceiverHP(d[0].name);
        }
        fetchHP();
    }, []);

    function handleChangeSenderHP(e) {
        setSelectedSenderHP([e.target.value]);
    }

    function handleChangeReceiverHP(e) {
        setSelectedReceiverHP([e.target.value]);
    }

    async function createConcent() {
        console.log("Value sender: " + selectedSenderHP);
        console.log("Value receiver: " + selectedReceiverHP);

        let senderId = await getHPByName(selectedSenderHP);
        let receiverId = await getHPByName(selectedReceiverHP);

        const res = await postConsent(senderId, receiverId);
    }

    return (
        <div className="ConsentForm">
            <form className="consent-form" onSubmit={handleSubmit} action="#">
                <p id="title">Create your consent</p>
                <label htmlFor="senderHP">
                    <span>Sender Healthcare Provider</span>
                    <select
                        name="senderHP"
                        onChange={(e) => {
                            handleChangeSenderHP(e);
                        }}
                    >
                        {senderHP.map((value, index) => (
                            <option key={index} name={value}>
                                {value}
                            </option>
                        ))}
                    </select>
                </label>
                <br></br>
                <label htmlFor="receiverHP">
                    <span>Reciever Healthcare Provider</span>
                    <select
                        name="receiverHP"
                        onChange={(e) => {
                            handleChangeReceiverHP(e);
                        }}
                    >
                        {receiverHP.map((value, index) => (
                            <option key={index} name={value}>
                                {value}
                            </option>
                        ))}
                    </select>
                </label>

                <div className="consent-agreement">
                    <label htmlFor="agreeConcent">
                        {" "}
                        I consent to share my medication record between the
                        healthcare providers
                        <input
                            type="checkbox"
                            name="agreeConcent"
                            value="agree"
                            required
                        />
                    </label>
                    <br></br>
                </div>

                {errorMsgVisibility && <p id="error-msg">{errorMsg}</p>}

                <input
                    id="submit-consent"
                    type="submit"
                    value="Create Consent"
                />
            </form>

            {consentCreated && (
                <div className="consent-created-div">
                    <p id="header">Your consent has been created!</p>
                    <p id="consent">{consentID}</p>
                </div>
            )}
        </div>
    );
};

export default ConsentForm;
