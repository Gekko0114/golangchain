package agent

const SYSTEM_MESSAGE_PREFIX = "Answer the following questions as best you can. You have access to the following tools:"
const FORMAT_INSTRUCTIONS = `The way you use the tools is by specifying a json blob.
Specifically, this JSON should have a action key (with the name of the tool to use) and a action_input key (with the input to the tool going here).

The only values that should be in the "action" field are: {{.ToolNames}}

The $JSON_BLOB should only contain a SINGLE action, do NOT return a list of multiple actions. Here is an example of a valid $JSON_BLOB:

  "Action_name": $TOOL_NAME,
  "Action_input": $INPUT

ALWAYS use the following format:

{
"Question": "the input question you must answer",
"Thought": "you should always think about what to do",
"Action":
{
  "Action_name": "$TOOL_NAME",
  "Action_input": "$INPUT"
}
}

Observation: the result of the action
... (this Thought/Action/Observation can repeat N times)
Thought: I now know the final answer
FinalAnswer: the final answer to the original input question`
const SYSTEM_MESSAGE_SUFFIX = "Begin! Reminder to always use the exact characters `FinalAnswer` when responding."
const HUMAN_MESSAGE = "{{ .Input }}\n\n{{ .Agent_scratchpad }}"
