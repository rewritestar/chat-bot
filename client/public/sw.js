// pwa
self.addEventListener('install', (e) => {
  console.log('SW installed');
});
self.addEventListener('activate', (e) => {
  console.log('SW activated');
});

//push
self.addEventListener('push', function (event) {
  const data = event.data.json();
  console.log(data);
  console.log(Notification.permission);
  event.waitUntil(
    self.registration.showNotification(data.title, {
      body: data.body,
      icon: '/icon-robot-192.png',
      badge: '/icon-robot-192.png',
      data: {
        url: data.url,
      },
    }),
  );
});

self.addEventListener('notificationclick', function (event) {
  event.notification.close();
  event.waitUntil(clients.openWindow(event.notification.data.url));
});
