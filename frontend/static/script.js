document.addEventListener('DOMContentLoaded', function() {
    const nameField = document.getElementById('editableName');
    const ageField = document.getElementById('editableAge');
    const backendUrl = 'http://localhost:3000';

    function getUser() {
        const userId = getCookie('userId');
        console.log("UserID:", userId);
        if (!userId) {
            nameField.innerText = "World";
            ageField.innerText = "0";
            return;
        };
        console.log("URL de requête:", `${backendUrl}/user/${userId}`);
        fetch(`${backendUrl}/user/${userId}`)
            .then(response => {
                if (!response.ok) {
                    console.error('Erreur de réponse:', response.status);
                    throw new Error('Réponse réseau non satisfaisante');
                }
                return response.json();
            })
            .then(data => {
                console.log('Données reçues:', data);
                nameField.innerText = data.name;
                ageField.innerText = data.age;
            })
            .catch(error => console.error('Error:', error));

    }

    function saveUser(name, age) {
        const userId = getCookie('userId');
        let url = `${backendUrl}/user`;
        let method = 'POST';
        let userData = { name, age };

        if (userId) {
            url = `${url}/${userId}`;
            method = 'PUT';
        }

        console.log('Données envoyées:', JSON.stringify(userData));
        fetch(url, {
            method: method,
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(userData),
        })
            .then(response => response.json())
            .then(data => {
                console.log('Success:', data);
                if (data.id && !userId) {
                    document.cookie = `userId=${data.id};path=/`;
                }
            })
            .catch((error) => {
                console.error('Error:', error);
            });
    }

    nameField.addEventListener('blur', function() {
        saveUser(nameField.innerText, ageField.innerText);
    });

    ageField.addEventListener('blur', function() {
        saveUser(nameField.innerText, ageField.innerText);
    });

    function getCookie(name) {
        let cookieValue = null;
        console.log("Cookies disponibles:", document.cookie); // Pour le débogage
        if (document.cookie && document.cookie !== '') {
            const cookies = document.cookie.split(';');
            for (let i = 0; i < cookies.length; i++) {
                const cookie = cookies[i].trim();
                if (cookie.substring(0, name.length + 1) === (name + '=')) {
                    cookieValue = decodeURIComponent(cookie.substring(name.length + 1));
                    break;
                }
            }
        }
        console.log("Valeur du cookie récupérée:", cookieValue); // Pour le débogage
        return cookieValue;
    }

    getUser();
});
