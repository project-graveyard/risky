---
description: Most enterprises use MVC pattern, so I want to experiment with it.
---

# ⚙️ Architecture

## Model View Controller (MVC)

### Model

**Responsibilities:**

* Retrieve data from the database.
* Save data to the database.
* Implement business logic.

**Interactions:**

* The model sends notifications to the view to update the display when data changes.

### View

**Responsibilities:**

* Render data from the model to the user interface.
* Send user input to the controller.

**Interaction:**

* The view listens for changes in the model and updates the presentation accordingly.

### Controller

**Responsibilities:**

* Interpret user inputs.
* Update the model based on user input.
* Decide which view to render.

**Interaction:**

* The controller manipulates the model and the view based on the input and application logic.

## Sample Workflow Interaction in MVC

1. The user interacts with the view (UI), such as clicking a button or entering text in a form.
2. The view sends the user action (event) to the controller.
3. The controller receives the input from the view. It processes the input, performs any necessary operations, and updates the model accordingly.
4. The model changes its state based on the operations performed by the controller (e.g., data retrieval, updates, or deletions).
5. The model sends notifications to the view about the changes in the data.
6. The view listens to the changes from the model. It updates the UI to reflect the latest state of the model.
