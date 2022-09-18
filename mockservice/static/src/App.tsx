import { useState, useEffect } from "react";
import axios from "axios";

function App() {
  const [isLogin, setIsLogin] = useState(false);
  useEffect(() => {
    //사용자가 로그인한 상태를 결정
  }, []);
  const handleLoginRequest = () => {
    console.log("버튼눌림");
    //통합 로그인 페이지 리디렉션해야 한다.
    //sso 서비스에 겟 요청을 보내어, 특정 페이지로 리디렉션 되도록 한다.
    axios.get("http://localhost:8080/loginpage").then((res) => {
      window.location.replace(res.request.responseURL);
    });
  };
  return (
    <div className="App">
      {isLogin ? (
        <>
          <section>관리자만 볼 수 있는 내용임</section>
          <section>로그인한 누구나 볼 수 있는 내용임</section>
        </>
      ) : (
        <button onClick={handleLoginRequest}>로그인 하러 가기</button>
      )}
    </div>
  );
}

export default App;
