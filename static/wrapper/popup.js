const channel = new BroadcastChannel('sw-messages');

const popup = new Popup({
    id: "ochanoco-popup",
    title: CONTENTS_TEXT.JA.title,
    content: CONTENTS_TEXT.JA.content,
});


channel.addEventListener('message', event => {
    console.log('Received', event.data);
    popup.show()

    setTimeout(() => {
        location.href = "/ochanoco/login"
    }, 5000)
});