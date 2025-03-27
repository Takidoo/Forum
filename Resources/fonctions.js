// on attend que le contenu DOM de la page soit chargé
document.addEventListener("DOMContentLoaded", function() { //on ajoute une écoute d'événement au chargement de la page
    console.log("DOM chargé"); //affiche un message console pour vérifie que le contenu est bien chargé
    
    document.getElementById("initialButtons").style.display="flex"; //on force l'affichage des boutons initiaux

    //on ajoute une écoute d'événement au clic sur le bouton Login
    document.getElementById("goToConnexion").addEventListener("click", function(e) {
        e.preventDefault(); //empêche le rechargement de la page
        console.log("click sur le bouton login"); //on affiche un message console pour vérifier que le clic fonctionne
        document.getElementById("initialButtons").style.display="none"; //cache les boutons initiaux à l'activation du bouton connexion
        document.getElementById("connexion").style.display="flex"; //rend la div connexion visible à l'activation du bouton connexion
    });

    //on ajoute une écoute d'événement au clic sur le bouton Register
    document.getElementById("goToInscription").addEventListener("click", function(e) {
        e.preventDefault(); //empêche le rechargement de la page
        console.log("click sur le bouton register"); //on affiche un message console pour vérifier que le clic fonctionne
        document.getElementById("initialButtons").style.display="none"; //rend les boutons invisibles à l'activation du bouton inscription
        document.getElementById("logo").style.display="none"; //rend le logo invisible
        document.getElementById("profilPicture").style.display="block"; //rend la div d'insertion de photo de profil visible
        document.getElementById("inscription").style.display="flex"; //rend la div inscription visible
    });

    //prévisualisation de l'image de profil
    document.getElementById("picture").addEventListener("change", function(event) { //affiche l'image sélectionnée dans la div preview à partir de l'input file
        const preview = document.getElementById("preview"); //recupère la div preview avec un id unique
        const file = event.target.files[0]; //recupère le fichier sélectionné à l'indice 1 dans le tableau files
    
        if (file) { //si un fichier est sélectionné
            const reader = new FileReader(); //on crée un objet FileReader pour lire le contenu du fichier que l'on stocke dans la variable reader
            reader.onload = function(e) { //lorsque le fichier est chargé
                preview.src = e.target.result; //affiche l'image sélectionnée
                preview.style.display = "block"; //rend l'image visible
            };
            reader.readAsDataURL(file); //lit le contenu du fichier
        }
    });

    //ajout de bouton back dans les div connexion et inscription
    //selection des conteneurs des formulaires de connexion et d'inscription
    const connexionDiv = document.getElementById("connexion"); //recupère le formulaire de connexion
    const inscriptionDiv = document.getElementById("inscription"); //recupère le formulaire d'inscription

    if (connexionDiv) { //si le formulaire de connexion existe
        console.log("formulaire de connexion trouvé, ajout du bouton Back"); //on affiche un message console
        
        //création de bouton retour au clic sur le bouton Submit du formulaire de connexion pour vérifier les identifiants et le rediriger vers la page d'accueil
        const backButtonConnexion = document.createElement("button"); //crée un bouton
        backButtonConnexion.textContent = "Back"; //ajoute le texte "Back" au bouton
        backButtonConnexion.type = "button"; //ajoute le type button au bouton
        backButtonConnexion.classList.add("backButton"); //ajoute la classe backButton au bouton

        //ajoute une écoute d'événement au clic sur le bouton retour
        backButtonConnexion.addEventListener("click", function(e) {
            e.preventDefault(); //on empêche le rechargement de la page
            console.log("click sur le bouton Back connexion"); //on affiche un message console pour vérifier que le clic du bouton retour fonctionne
            connexionDiv.style.display="none"; //on rend le formulaire invisible
            document.getElementById("initialButtons").style.display="flex"; //les boutons initiaux visibles
        });

        //on ajoute le bouton au formulaire de connexion
        connexionDiv.appendChild(backButtonConnexion) //ajoute le bouton à la div du formulaire de connexion
        }else{
            console.error("formulaire de connexion non trouvé"); //on affiche un message d'erreur sur la console si la div du formulaire de connexion n'est pas trouvé
        }

    if (inscriptionDiv) { //si le formulaire d'inscription existe
        console.log("formulaire d'inscription trouvé, ajout du bouton Back"); //on affiche un message console
        
        //on créer de bouton retour au clic sur le bouton Submit du formulaire d'inscription pour vérifier les identifiants et le rediriger vers la page d'accueil
        const backButtonInscription = document.createElement("button"); //on crée un bouton
        backButtonInscription.textContent = "Back"; //on ajoute le texte "Back" au bouton
        backButtonInscription.type = "button"; //on ajoute le type button au bouton
        backButtonInscription.classList.add("backButton"); //on ajoute la classe backButton au bouton

        //on ajoute une écoute d'événement au clic sur le bouton retour
        backButtonInscription.addEventListener("click", function(e) {
            e.preventDefault(); //on empêche le rechargement de la page
            console.log("click sur le bouton Back connexion"); //on affiche un message console pour vérifier que le clic du bouton retour fonctionne
            inscriptionDiv.style.display="none"; //on rend le formulaire invisible
            document.getElementById("initialButtons").style.display="flex"; //les boutons initiaux sont visibles
        });

        //on ajoute le bouton au formulaire d'inscription
        inscriptionDiv.appendChild(backButtonInscription) //on ajoute le bouton à la div du formulaire d'inscription
        }else{
            console.error("formulaire d'inscription non trouvé"); //on affiche un message d'erreur sur la console si la div du formulaire d'inscription n'est pas trouvé
        }

    // Get the form by its ID
    const loginForm = document.getElementById('loginForm');

    // Only if the form exists (it should), add a 'submit' listener
    if (loginForm) {
    loginForm.addEventListener('submit', function(e) {
        e.preventDefault(); // Stop the page from reloading

        // Collect the form data
        const formData = new FormData(loginForm);

        // Send it to the Go handler using fetch
        fetch('/api/login', {
        method: 'POST',
        body: formData
        })
        .then(response => {
        // Even if the response is 401 or 400, fetch won't throw an error,
        // so we can still read JSON. But let's handle OK vs. error below:
        return response.json().then(data => ({ status: response.status, body: data }));
        })
        .then(({ status, body }) => {
        // The Go handler returns either {"success": "..."} or {"erreur": "..."}
        if (status === 200 && body.success) {
            // Show success message (popup or alert)
            alert("Succès : " + body.success);
            // Optionally redirect somewhere:
            // window.location.href = "/someOtherPage";
        } else {
            // Show error popup
            // body.erreur might be "Identifiants invalides" or some other message
            alert("Erreur : " + body.erreur);
        }
        })
        .catch(error => {
        console.error("Fetch error:", error);
        alert("Une erreur inattendue est survenue.");
        });
    });
    }
});

function showPopup(message) {
    const popup = document.getElementById("popup");
    const messageElem = document.getElementById("popupMessage");
    messageElem.textContent = message;
    popup.style.display = "block";
  }

function closePopup() {
    document.getElementById("popup").style.display = "none";
}