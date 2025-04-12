/*ATTENTE DU CHARGEMENT INTEGRALE DE LA PAGE*/
document.addEventListener("DOMContentLoaded", function() { //on ajoute une écoute d'événement au chargement de la page
    console.log("DOM chargé"); //affiche un message console pour vérifie que le contenu est bien chargé
    
    /*GESTION DE L'AFFICHAGE DU MENU DEROULANT D'AUTHENTIFICATION*/
    const authBtn = document.getElementById("authBtn"); //on récupère le bouton d'authentification
    const authDropdown = document.getElementById("authDropdown"); //on récupère le menu déroulant d'authentification

    if(authBtn && authDropdown) { //si le bouton et le menu existent
        authBtn.addEventListener("click", function(e) { //on ajoute une écoute d'événement au clic sur le bouton d'authentification
            e.stopPropagation(); //on empêche la propagation de l'événement

            if(authDropdown.style.display === "block") { //si le menu est déjà affiché
                authDropdown.style.display = "none"; //on le cache
                authBtn.classList.remove("active"); //on enlève la classe active du bouton
            }else{
                authDropdown.style.display = "block"; //sinon, on l'affiche
                authBtn.classList.add("active"); //on ajoute la classe active au bouton
            }
        });

        /*ferme le menu si on clic ailleurs*/
        document.addEventListener("click", function(event) { //on ajoute une écoute d'événement au clic sur le document
            if (!authBtn.contains(event.target) && !authDropdown.contains(event.target)) { //si le clic n'est pas sur le bouton 
                authDropdown.style.display = "none"; //on cache le menu
                authBtn.classList.remove("active"); //on enlève la classe active du bouton
            }
        });
    }

    /*ON AFFICHE LES BOUTONS INITIAUX*/
    document.getElementById("initialButtons").style.display="flex"; //on force l'affichage des boutons initiaux

    /*GESTIONNAIRES D'EVENEMENTS POUR LE MENU DEROULANT ET LES BOUTONS MOBILES*/
    

    /*AFFICHER OU CACHER LE MENU AU CLIC DU BOUTON AUTHENTICATION*/
    authBtn.addEventListener("click", function(e) { //on ajoute une écoute d'événement au clic sur le bouton d'authentification
        e.stopPropagation(); 
        if(authDropdown.style.display === "block") { //si le menu est déjà affiché
            authDropdown.style.display = "none"; //on le cache
            authBtn.classList.remove("active"); //on enlève la classe active du bouton
        }else{
            authDropdown.style.display = "block"; //sinon, on l'affiche
            authBtn.classList.add("active"); //on ajoute la classe active au bouton
        }
    });

    /*FERMER LE MENU AU CLIC EN DEHORS SUR LA PAGE*/
    document.addEventListener("click", function(event) { //on ajoute une écoute d'événement au clic sur le document
        if (!authBtn.contains(event.target) && !authDropdown.contains(event.target)) { //si le clic n'est pas sur le bouton ou le menu
            authBtn.classList.remove("active"); //on enlève la classe active du bouton
            authDropdown.style.display = "none"; //on cache le menu
        }
    });

    /*AFFICHAGE INITIALE DES BOUTONS*/
    document.getElementById("initialButtons").style.display="flex"; //on force l'affichage des boutons initiaux

    /*AFFICHAGE DYNAMIQUE AU CLIC DU BOUTON LOGIN*/
    document.getElementById("goToConnexion").addEventListener("click", function(e) {//on ajoute une écoute d'événement au clic sur le bouton Login
        e.preventDefault(); //empêche le rechargement de la page
        console.log("click sur le bouton login"); //on affiche un message console pour vérifier que le clic fonctionne
        authDropdown.style.display = "none"; //cache dropdow, le menu d'authentification
        authBtn.classList.remove("active"); //enlève la classe active du bouton

        document.getElementById("logo").style.display="block"; //affiche le logo
        document.getElementById("connexion").style.display="flex"; //rend la div connexion visible à l'activation du bouton connexion
    });

    /*AFFICHAGE DYNAMIQUE AU CLIC DU BOUTON REGISTER*/
    document.getElementById("goToInscription").addEventListener("click", function(e) {//on ajoute une écoute d'événement au clic sur le bouton Register
        e.preventDefault(); //empêche le rechargement de la page
        console.log("click sur le bouton register"); //on affiche un message console pour vérifier que le clic fonctionne
        authDropdown.style.display = "none"; //cache le menu d'authentification
        authBtn.classList.remove("active"); //enlève la classe active du bouton

        document.getElementById("logo").style.display="none"; //rend le logo invisible
        document.getElementById("profilPicture").style.display="block"; //rend la div d'insertion de photo de profil visible
        document.getElementById("inscription").style.display="flex"; //rend la div inscription visible
    });

    /*ADAPTER LES LIENS DU MENU DEROULANT POUR LES OUBLIS D'ID ET DE MOT DE PASSE*/
    document.getElementById("forgetIdlink").addEventListener("click", function(e) { //on ajoute une écoute d'événement au clic sur le lien d'oubli d'id
        e.preventDefault(); //empêche le rechargement de la page
        console.log("click sur le lien d'oubli d'id"); //on affiche un message console pour vérifier que le clic fonctionne
        authDropdown.style.display = "none"; //cache le menu d'authentification
        authBtn.classList.remove("active"); //enlève la classe active du bouton
        showUpdateSection('username'); //on affiche la section de mise à jour   
    });

    document.getElementById("forgetPasswordlink").addEventListener("click", function(e) { //on ajoute une écoute d'événement au clic sur le lien d'oubli de mot de passe
        e.preventDefault(); //empêche le rechargement de la page
        console.log("click sur le lien d'oubli de mot de passe"); //on affiche un message console pour vérifier que le clic fonctionne
        authDropdown.style.display = "none"; //cache le menu d'authentification
        authBtn.classList.remove("active"); //enlève la classe active du bouton
        showUpdateSection('password'); //on affiche la section de mise à jour
    });

    /*FONCTION POUR GERER LA RESPONSIVE AU CHARGEMENT DE LA PAGE*/
    function handleResponsiveDisplay(){
        const mediaQuery = window.matchMedia("(max-width: 48em)"); //on crée une media query pour les écrans de moins de 768px

        if (mediaQuery.matches) { //si la media query est vérifiée
            if(document.getElementById("initialButtons")) { //si on est en mode desktop ou tablette, on cache les éléments mobiles
                document.getElementById("initialButtons").style.display = "none";
            }
        }else{
            if(document.getElementById("manNav")) { //si on est en mode mobile, on affiche les éléments desktop
                document.getElementById("mainNav").style.display = "none";
            }
            if(document.getElementById("initialButtons")) { //si on est en mode desktop ou tablette, on affiche les éléments mobiles
                document.getElementById("initialButtons").style.display = "flex";
            }
        }
    }

    /*PREVISUALISATION DE L'IMAGE DE PROFIL*/
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

    /*FONCTION QUI PERMET DE REINITIALISER LES FORMULAIRES ET LES IMAGES*/
    function resetElements(){
        //on réintialise le formulaire d'inscription
        if (document.getElementById("inscription").querySelector("form")){
            document.getElementById("inscription").querySelector("form").reset();
        }
        //on réintialise le formulaie de connexion
        if (document.getElementById("connexion").querySelector("form")){
            document.getElementById("connexion").querySelector("form").reset();
        }
        //on réinitialise l'image de profil
        document.getElementById("preview").style.display="none"; //cache l'élément preview
        document.getElementById("preview").src=""; //vide la source de l'image
            if (document.getElementById("picture")){ //réinitialise le champ input de type file (l'image posté)
                document.getElementById("picture").value="";
            }
    }


    /*AJOUT DE BOUTONS BACK DANS LES DIV CONNEXION ET INSCRIPTION*/
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
            document.getElementById("logo").style.display="block"; //affiche à nouveau le logo
            document.getElementById("profilPicture").style.display="none"; //fait bien disparaitre la section avec la photo de profil
            document.getElementById("initialButtons").style.display="flex"; //les boutons initiaux visibles
            resetElements(); //on fait appel à la fonction qui réinitialise les formulaires et l'image
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
            console.log("click sur le bouton Back inscription"); //on affiche un message console pour vérifier que le clic du bouton retour fonctionne
            inscriptionDiv.style.display="none"; //on rend le formulaire invisible
            document.getElementById("profilPicture").style.display="none"; //on cache la section de photo de profil
            document.getElementById("logo").style.display="block"; //on affiche le logo à la place de la photo de profil
            document.getElementById("initialButtons").style.display="flex"; //les boutons initiaux sont visibles
            resetElements(); //on fait appel à la fonction qui réinitialise les formulaires et l'image
        });

        //on ajoute le bouton au formulaire d'inscription
        inscriptionDiv.appendChild(backButtonInscription) //on ajoute le bouton à la div du formulaire d'inscription
        }else{
            console.error("formulaire d'inscription non trouvé"); //on affiche un message d'erreur sur la console si la div du formulaire d'inscription n'est pas trouvé
        }

    /*FONCTIONNALITE DE MISE A JOUR*/
    //création de lien pour accèder à la section de mise à jour
    const forgetIdLink = document.getElementById("forgetIdlink");
    const forgetPasswordLink = document.getElementById("forgetPasswordlink");

    //mise à jour des liens
    forgetIdLink.addEventListener("click", function(e){
        e.preventDefault();
        showUpdateSection('username');
    });

    forgetPasswordLink.addEventListener("click", function(e){
        e.preventDefault();
        showUpdateSection('password');
    });

    //affiche la section de modification
    function showUpdateSection(updateType){
        //on cache les éléments suivant:
        document.getElementById("initialButtons").style.display="none";
        document.getElementById("connexion").style.display="none";
        document.getElementById("inscription").style.display="none";
        document.getElementById("profilPicture").style.display="none";
        document.getElementById("updates").style.display="none";
        //on montre les éléments suivant:
        document.getElementById("logo").style.display="block";
        document.getElementById("update").style.display="flex";
        // on cache les deux sections par défaut:
        document.getElementById("usernameUpdateData").style.display="none";
        document.getElementById("passwordUpdateData").style.display="none";
        document.getElementById("updateSubmit").style.display="block";

        //on selectionne quel type de mise à jour on veut faire (mot de passe ou id)
        if (updateType === 'username'){
            document.getElementById("usernameUpdateData").style.display="block";
            document.getElementById("passwordUpdateData").style.display="none";
        }else if(updateType === 'password'){
            document.getElementById("passwordUpdateData").style.display="block";
            document.getElementById("usernameUpdateData").style.display="none";
        }
    }
    //on ajoute des écoutes d'évenements pour les boutons de mises à jour
    document.getElementById("showUsernameUpdate").addEventListener("click", function(){
        //on montre les éléments suivant:
        document.getElementById("usernameUpdateData").style.display="block";
        document.getElementById("updateSubmit").style.display="block";
        //on cache les éléments suivant:
        document.getElementById("passwordUpdateData").style.display="none";
    });

    document.getElementById("showPasswordUpdate").addEventListener("click", function(){
        //on montre les éléments suivant:
        document.getElementById("passwordUpdateData").style.display="block";
        document.getElementById("updateSubmit").style.display="block";
        //on cache les éléments suivant:
        document.getElementById("usernameUpdateData").style.display="none";
    });

    //on créer une fonction pour gérer les messages d'erreur sur l'interface utilisateur
    function showError(inputElement, message){
        if(inputElement){ //on vérifie si l'élément existe et s'il contient une saisie
            inputElement.textContent = message;
            inputElement.style.display = "block";
        }else{
            alert(message); //si ce n'est pas le cas, on affiche une message d'erreur
        }
    }
    //on valide le formulaire de modifications
    const updateForm = document.querySelector("#update form");
    if(updateForm){
        updateForm.addEventListener("submit", function(e){
        e.preventDefault();

        //on vérifie l'email de l'utilisateur et on rend le champ obligatoire
        const userEmail = document.getElementById("userEmail").value.trim();
        if(!userEmail){
            alert("Please enter your email"); 
            return;
        }

        //Vérification de la validité de l'email
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/; //expression régulière pour vérifier la validité de l'email
        if (!emailRegex.test(userEmail)) { //si l'email ne correspond pas à l'expression régulière
            alert("Please enter a valid email address"); //on déclenche une alerte
            return;
        }
    
        //on vérifie le nom de l'utilisateur et on rend le champ obligatoire
        const currentUsername = document.getElementById("currentUsername").value.trim();
        if(!currentUsername){
            alert("Please enter your current username");
            return;
        }

        //on vérifie que l'utilisateur a bien sélectionné une question de sécurité
        const securityQuestion = document.getElementById("securityAnswer").value.trim();
        if (!securityQuestion || securityQuestion === ""){
            alert("Please select a security question");
            return;
        }

        //si l'utilisateur n'a pas répondu à la question de sécurité:
        const securityAnswer = document.getElementById("securityAnswer").value.trim();
        if (!securityAnswer){ //si l'entrée utilisateur n'a pas été saisie
            alert("Please answer the security question"); //on déclanche une alerte en affichant un message sur l'interface utilisateur
            return;
        }
        //si l'utilisateur n'a pas entré un nouvel identifiant:
        if (document.getElementById("usernameUpdateData").style.display === "block"){ //on verifie que l'input du nouvel id est visible 
            const newUsername = document.getElementById("newUsername").value.trim(); //si c'est le cas, on récupère la valeur du champ de texte en enlevant les espaces inutiles au début et à la fin
            if (!newUsername){
                alert("Please enter a new username");
                return;
            }
        }
        
        //si l'utilisateur n'a pas entré un nouveau mot de passe:
        if (document.getElementById("passwordUpdateData").style.display === "block"){
            const newPassword = document.getElementById("newPassword").value.trim(); //on vérifie les deux saisies
            const confirmPassword = document.getElementById("confirmPassword").value.trim();

            if(!newPassword){
                alert("Please enter a new password");
                return;
            }
            if(!confirmPassword){
                alert("Please enter a new password");
                return;
            }
            if(newPassword !== confirmPassword){
                alert("Password do not match");
                return;
            }
        }
    
        //si les modifications on bien été mise à jour, on affiche un message d'alerte:
        alert("Your information has been successfully updated");

        //on réinitialise l'affichage initiale
        updateDiv.style.display="none"; //on cache la div de mise à jour
        document.getElementById("logo").style.display="block"; //on affiche le logo et les boutons initiaux
        document.getElementById("initialButtons").style.display="flex";
        document.getElementById("updates").style.display="flex";
        updateForm.reset();//on efface toutes les saisies utilisateurs
        });
    }else{
        console.error("formulaire de modification est introuvable")
    }

    /*AJOUT D'UN BOUTON BACK POUR LA MISE A JOUR UTILISATEUR*/
    const updateDiv = document.getElementById("update");
    if (updateDiv){
        const backButtonUpdate = document.createElement("button");
        backButtonUpdate.textContent = "Back";
        backButtonUpdate.type = "button";
        backButtonUpdate.classList.add("backButton");

        backButtonUpdate.addEventListener("click", function(e){
            e.preventDefault(); // empêche le refresh
            updateDiv.style.display ="none"; //cache la section de modification
            
            document.getElementById("usernameUpdateData").style.display="none"; //réinitialise les champs de saisie
            document.getElementById("passwordUpdateData").style.display="none";
            document.getElementById("updateSubmit").style.display="none";

            document.getElementById("initialButtons").style.display = "flex"; //réinitialise l'affichage initiale
            document.getElementById("logo").style.display="block";
            document.getElementById("updates").style.display="flex";

            if (updateForm){ //réinitialise le formulaire
                updateForm.reset();
            }
        });
        updateDiv.appendChild(backButtonUpdate);
    }
});

/*S'EXECTUTE A CHAQUE FOIS QUE LA PAGE EST REDIMENSIONNEE*/
window.addEventListener("resize", handleResponsiveDisplay); //on ajoute une écoute d'événement au redimensionnement de la page
handleResponsiveDisplay(); //on appelle la fonction pour gérer la responsive
window.addEventListener("load", handleResponsiveDisplay); //on appelle la fonction pour gérer la responsive

//fonction pour afficher/masquer le menu deroulant
function Filter() {
    document.getElementById("myDropdown").classList.toggle("show");
}

//affiche les discussions
document.addEventListener("DOMContentLoaded", function () {
    // Appel à l'API pour récupérer les discussions
    fetch('/API/FetchThreadPosts?limit=6') // Ajoutez un paramètre `limit=6` pour limiter à 6 discussions
        .then(response => { // Traite la réponse de l'API
            if (!response.ok) { // Vérifie si la réponse est correcte
                throw new Error("Erreur lors de la récupération des discussions"); // Lance une erreur si la réponse n'est pas correcte
            }
            return response.json(); // Convertit la réponse en JSON
        })
        .then(data => { // Traite les données JSON
            const discussionsContainer = document.getElementById("discussions"); // Récupère le conteneur des discussions
            discussionsContainer.innerHTML = ""; // Vide le conteneur

            if (data.length === 0) { // Vérifie si aucune discussion n'est disponible
                discussionsContainer.innerHTML = "<li>No discussions available.</li>"; // Affiche un message si aucune discussion n'est trouvée
                return;
            }

            // Ajoute chaque discussion au conteneur
            data.forEach(post => { // Parcourt chaque discussion
                const li = document.createElement("li"); // Crée un nouvel élément de liste
                li.innerHTML = `<a href="/post/${post.PostID}">${post.Content}</a>`; // Définit le contenu de l'élément de liste avec un lien vers la discussion
                discussionsContainer.appendChild(li); // Ajoute l'élément de liste au conteneur
            });
        })
        .catch(error => { // Gère les erreurs
            console.error(error); // Affiche l'erreur dans la console
            document.getElementById("discussions").innerHTML = "<li>Error loading discussions.</li>"; // Affiche un message d'erreur si la récupération échoue
        });
});

//discussions likées
document.addEventListener("DOMContentLoaded", function () {
    // Appel à l'API pour récupérer les discussions likées
    fetch('/API/LikeThread?limit=6') // Remplacez par l'URL correcte de votre API
        .then(response => {
            if (!response.ok) {
                throw new Error("Erreur lors de la récupération des discussions likées");
            }
            return response.json();
        })
        .then(data => {
            const likedDiscussionsContainer = document.getElementById("liked-discussions");
            const likeSection = document.querySelector(".like");

            likedDiscussionsContainer.innerHTML = ""; // Vide le conteneur

            if (data.length === 0) {
                // Si aucune discussion n'est likée, masque la section
                likeSection.style.display = "none";
                return;
            }

            // Ajoute chaque discussion likée au conteneur
            data.forEach(thread => {
                const li = document.createElement("li");
                li.innerHTML = `<a href="/thread/${thread.ThreadID}">${thread.Title}</a>`;
                likedDiscussionsContainer.appendChild(li);
            });
        })
        .catch(error => {
            console.error(error);
            document.getElementById("liked-discussions").innerHTML = "<li>Error loading liked discussions.</li>";
        });
});