import {useLoginStore} from 'stores/login-store'
import { jwtDecode } from 'jwt-decode'
import { Notify } from 'quasar'

export default {
  name: 'MyLayout',
  components:{},

  setup () {
    const store = useLoginStore();

    const checkClickSession = () => {
      const refreshToken = store.getRefreshToken.value || '';
      if(!refreshToken) return;
      const decodedToken = jwtDecode(refreshToken);
      const currentTime = Date.now() / 1000;
      const expirationTime = decodedToken.exp || 0;
      const timeRemaining = expirationTime - currentTime;
      console.log( '>>>>>>>>>>>>>>>>', timeRemaining);
      if(timeRemaining < 0 ){
        store.setToken();
        Notify.create({message:'Session expirada, por favor vuelva a ingresar.',color:'red'})
      }
    }

    return {
      checkClickSession,
    }
  }
}
