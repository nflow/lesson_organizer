<template>
  <div class="tw-w-full">
    <div class="tw-bg-white">
      <Draggable
        :id="boardId ? boardId : methodId"
        :data-parent-type="boardId ? 'board' : 'method'"
        v-model="contents"
        @end="onMoveContent"
        tag="table"
        animation="150"
        group="ideas"
        item-key="id"
        class="
          tw-min-w-max
          tw-w-full
          tw-table-auto
          tw-text-gray-600
          tw-text-sm
          tw-font-light
        "
      >
        <template #item="{ element }">
          <tr
            :id="element.id"
            class="tw-border-b tw-border-gray-200 hover:tw-bg-gray-100"
          >
            <td class="tw-py-3 tw-px-6 tw-text-left">
              <div class="tw-relative">
                <span>{{ element.text }}</span>
                <div
                  class="
                    tw-absolute tw-inset-y-0 tw-right-0 tw-flex tw-items-center
                  "
                >
                  <span
                    class="tw-cursor-pointer hover:tw-text-gray-200"
                    @click="onModify(element)"
                  >
                    <svg
                      class="tw-w-4 tw-h-full"
                      fill="none"
                      stroke="currentColor"
                      viewBox="0 0 24 24"
                      xmlns="http://www.w3.org/2000/svg"
                    >
                      <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"
                      ></path>
                    </svg>
                  </span>
                  <span
                    class="tw-cursor-pointer hover:tw-text-red-400"
                    @click="onRemove(element)"
                  >
                    <svg
                      class="tw-w-4 tw-h-full"
                      fill="none"
                      stroke="currentColor"
                      viewBox="0 0 24 24"
                      xmlns="http://www.w3.org/2000/svg"
                    >
                      <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
                      ></path>
                    </svg>
                  </span>
                </div>
              </div>
            </td>
          </tr>
        </template>
        <template #footer>
          <tr
            class="
              tw-border-b tw-py-3 tw-px-6 tw-text-left tw-border-gray-200
              hover:tw-bg-gray-100
            "
          >
            <td>
              <div class="tw-relative">
                <div
                  class="
                    tw-absolute
                    tw-inset-y-0
                    tw-right-0
                    tw-pl-2
                    tw-flex
                    tw-items-center
                    tw-pointer-events-none
                  "
                >
                  <svg
                    class="tw-w-6 tw-h-6 tw-text-gray-300"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                    xmlns="http://www.w3.org/2000/svg"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M12 9v3m0 0v3m0-3h3m-3 0H9m12 0a9 9 0 11-18 0 9 9 0 0118 0z"
                    />
                  </svg>
                </div>
                <input
                  id="new-entry"
                  v-model="newEntryInput"
                  type="text"
                  name="new-entry"
                  class="
                    tw-inset-0
                    tw-py-3
                    tw-px-6
                    tw-block
                    tw-w-full
                    tw-pl-9
                    tw-pr-12
                    tw-border-gray-200
                    tw-hover:bg-gray-100
                  "
                  placeholder="New amazing content ..."
                  @keydown.enter="addNew"
                />
              </div>
            </td>
          </tr>
        </template>
      </Draggable>
    </div>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, ref } from "vue";
import { ContentDto, CreateContentDto } from "@/types/content";
import Draggable from "vuedraggable";
import {
  deleteContent,
  getBoardContents,
  getMethodContents,
  postBoardContent,
  postMethodContent,
  putContent,
} from "@/api/board";
import { useQueryClient } from "vue-query";

export default defineComponent({
  name: "ContentList",
  components: {
    Draggable,
  },
  props: {
    boardId: {
      type: String,
      required: false,
      default: undefined,
    },
    methodId: {
      type: String,
      required: false,
      default: undefined,
    },
  },
  setup(props) {
    const queryClient = useQueryClient();

    const contentsQuery = props.methodId
      ? getMethodContents(props.methodId)
      : getBoardContents(ref(props.boardId));

    const contents = computed({
      get: () => {
        return contentsQuery.data.value;
      },
      set: (data) => {
        queryClient.setQueryData(
          props.methodId
            ? ["method_contents", props.methodId]
            : ["board_contents", props.boardId],
          data
        );
      },
    });

    const createContent = props.methodId
      ? postMethodContent(props.methodId)
      : postBoardContent(ref(props.boardId));

    const newEntryInput = ref("");
    const addNew = (): void => {
      if (!newEntryInput.value) {
        return;
      }
      const newEntry: CreateContentDto = {
        text: newEntryInput.value,
      };

      createContent.mutate(newEntry, {
        onSuccess: () => {
          queryClient.invalidateQueries(
            props.methodId
              ? ["method_contents", props.methodId]
              : ["board_contents", props.boardId]
          );
          newEntryInput.value = "";
        },
      });

      const newList: ContentDto[] = [...contentsQuery.data.value, newEntry];

      queryClient.setQueryData(
        props.methodId
          ? ["method_contents", props.methodId]
          : ["board_contents", props.boardId],
        newList
      );
    };

    const updateContent = putContent();

    if (!updateContent) {
      return;
    }

    const onMoveContent = (evt: any) => {
      console.log(evt);
      if (
        evt.newDraggableIndex == evt.oldDraggableIndex &&
        evt.to.id == evt.from.id
      ) {
        return;
      }

      const isBoard = (object: HTMLElement): boolean => {
        return object.attributes.getNamedItem("data-parent-type")?.value ==
          "board"
          ? true
          : false;
      };

      let afterId = undefined;
      if (evt.newDraggableIndex > 0) {
        const contents = queryClient.getQueryData<ContentDto[]>([
          isBoard(evt.to) ? "board_contents" : "method_contents",
          evt.to.id,
        ]);
        if (contents) {
          afterId = contents[evt.newDraggableIndex - 1].id;
        }
      }

      updateContent.mutate(
        {
          contentId: evt.item.id,
          parentBoardId: isBoard(evt.to) ? evt.to.id : undefined,
          parentMethodId: !isBoard(evt.to) ? evt.to.id : undefined,
          afterContentId: afterId,
        },
        {
          onSuccess: () => {
            // TODO: Could return the new list of phases .. or maybe not.
            queryClient.invalidateQueries([
              isBoard(evt.to) ? "board_contents" : "method_contents",
              evt.to.id,
            ]);
            if (evt.from.id != evt.to.id) {
              queryClient.invalidateQueries([
                isBoard(evt.from) ? "board_contents" : "method_contents",
                evt.from.id,
              ]);
            }
          },
        }
      );
    };

    const removeContent = deleteContent();

    if (!removeContent) {
      return;
    }

    const onRemove = (content: ContentDto) => {
      removeContent.mutate(content.id, {
        onSuccess: () => {
          queryClient.invalidateQueries(
            props.methodId
              ? ["method_contents", props.methodId]
              : ["board_contents", props.boardId]
          );
        },
      });

      const newList = [...contentsQuery.data.value];

      newList.splice(newList.indexOf(content), 1);

      queryClient.setQueryData(
        props.methodId
          ? ["method_contents", props.methodId]
          : ["board_contents", props.boardId],
        newList
      );
    };

    return {
      newEntryInput,
      onMoveContent,
      contents,
      addNew,
      onRemove,
    };
  },
});
</script>
