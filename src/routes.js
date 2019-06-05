import Home from './components/Home.vue';
import Team from './components/Team.vue';
import Salon from './components/Salon.vue';
import Reviews from './components/Reviews.vue';
import Contact from './components/Contact.vue';


export const routes = [
    { path: '', component: Home },
    { path: '/team', component: Team },
    { path: '/salon', component: Salon },
    { path: '/reviews', component: Reviews },
    { path: '/contact', component: Contact }
]