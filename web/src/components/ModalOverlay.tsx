"use client";

import { withModal } from "./HOC";
import { NewPostForm } from "./NewPostForm";

export const ModalOverlay = withModal(NewPostForm);
