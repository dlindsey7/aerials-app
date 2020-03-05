import React from "react";
import "./App.css";
// import the Container Component from the semantic-ui-react
import { Container } from "semantic-ui-react";
// import the ToDoList component
import SkillList from "./Skill-List";
function App() {
  return (
    <div>
      <Container>
        <SkillList />
      </Container>
    </div>
  );
}
export default App;
