document.addEventListener('DOMContentLoaded', () => {
    const gamesHistory = document.getElementById('historyList');
    const chatContent = document.getElementById('chatContent');
    const matchTitle = document.getElementById('matchTitle');
    const liveScore = document.getElementById('liveScore');
    const additionalDetails = document.getElementById('detailsContent');

    let currentMatchId = null;

    function initializePage() {
        fetch('http://localhost:2030/matches')
            .then(response => response.json())
            .then(data => {
                if (data && data.matches && data.matches.length > 0) {
                    data.matches.forEach((match, index) => {
                        if (match.team1 && match.team2) { 
                            const li = document.createElement('li');
                            li.innerHTML = `<strong>${match.team1.name}</strong> vs <strong>${match.team2.name}</strong>`;
                            li.addEventListener('click', () => loadGameDetails(match.id));
                            gamesHistory.appendChild(li);

                            if (index === 0) {
                                loadGameDetails(match.id);
                            }
                        }
                    });
                } else {
                    displayNoActiveMatches();
                }
            })
            .catch(error => {
                console.error('Error fetching matches:', error);
                displayNoActiveMatches();
            });
    }

    function displayNoActiveMatches() {
        gamesHistory.innerHTML = "<p>No active matches</p>";
        matchTitle.textContent = "No Active Matches";
        liveScore.textContent = "";
        chatContent.innerHTML = "<p>There are no active matches at the moment.</p>";
        additionalDetails.innerHTML = "<p>Please check back later for upcoming matches.</p>";
    }

    function loadGameDetails(matchID) {
        currentMatchId = matchID;
        fetch(`http://localhost:2030/match/${matchID}`)
            .then(response => response.json())
            .then(data => {
                const match = data.match;
                const messages = data.messages;

                if (match && match.team1 && match.team2) {
                    updateMatchHeader(match);
                    chatContent.innerHTML = '';

                    if (messages.messages && messages.messages.length > 0) {
                        messages.messages.forEach(msg => appendMessage(msg));
                    } else {
                        appendMessage({content: "No updates yet", sent_at: new Date()});
                    }

                    updateAdditionalDetails(match);

                    // Close existing WebSocket connection if any
                    if (window.matchSocket) {
                        window.matchSocket.close();
                    }

                    // Open a new WebSocket connection for this match
                    window.matchSocket = new WebSocket('ws://localhost:2030/ws');
                    window.matchSocket.onmessage = function(event) {
                        const msg = JSON.parse(event.data);
                        if (msg.match_id === currentMatchId) {
                            appendMessage(msg);
                            updateScore(msg);
                        }
                    };
                }
            })
            .catch(error => {
                console.error('Error fetching match details:', error);
                displayNoActiveMatches();
            });
    }

    function updateMatchHeader(match) {
        matchTitle.textContent = `${match.team1.name} vs ${match.team2.name}`;
        liveScore.textContent = `${match.result1} - ${match.result2}`;
    }

    function updateAdditionalDetails(match) {
        additionalDetails.innerHTML = `
            <p><strong>Date:</strong> ${new Date(match.date).toLocaleString()}</p>
            <p><strong>Status:</strong> ${match.status}</p>
            <p><strong>Round:</strong> ${match.round}</p>
            <h3>${match.team1.name}</h3>
            <p><strong>Country:</strong> ${match.team1.country}</p>
            <p><strong>Group:</strong> ${match.group_name}</p>
            <h3>${match.team2.name}</h3>
            <p><strong>Country:</strong> ${match.team2.country}</p>
            <p><strong>Group:</strong> ${match.group_name}</p>
        `;

    }

    function appendMessage(msg) {
        const messageDiv = document.createElement('div');
        messageDiv.className = 'message';
        
        const content = document.createElement('p');
        content.textContent = msg.content;
        messageDiv.appendChild(content);

        const time = document.createElement('p');
        time.className = 'message-time';
        time.textContent = new Date(msg.sent_at).toLocaleTimeString();
        messageDiv.appendChild(time);

        chatContent.insertBefore(messageDiv, chatContent.firstChild);
    }

    function updateScore(msg) {
        if (msg.score) {
            liveScore.textContent = msg.score;
            liveScore.classList.add('score-update');
            setTimeout(() => {
                liveScore.classList.remove('score-update');
            }, 500);
        }
    }

    initializePage();
});
