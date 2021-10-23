import { Ref } from "@vue/runtime-core";

type Components = {
  [key: string]: Ref;
};

type Rule = {
  (val: string | undefined): void;
};

type Rules = {
  [key: string]: Rule[];
};

export interface GenericDialog<ModelType> {
  visible: Ref<boolean>;
  model: Ref<ModelType>;
  components?: Components;
  rules?: Rules;
  onOpen: (model: ModelType) => void;
  onSubmit: VoidFunction;
  onCancle: VoidFunction;
}
