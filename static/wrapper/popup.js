const channel = new BroadcastChannel('sw-messages');

const popup = new Popup({
    id: "torima-popup",
    title: CONTENTS_TEXT.JA.title,
    content: CONTENTS_TEXT.JA.content,
});


channel.addEventListener('message', event => {
    console.log('Received', event.data);
    popup.show()

    setTimeout(() => {
        location.href = "/torima/login"
    }, 5000)
});