"use client";
import React, { ComponentType, useEffect } from "react";

interface ModalProps {
  isOpen: boolean;
  onClose: () => void;
}

export function withModal<T extends object>(
  WrappedComponent: ComponentType<T>
) {
  const ModalComponent = (props: T & ModalProps) => {
    const { isOpen, onClose, ...rest } = props;

    useEffect(() => {
      if (isOpen) {
        document.body.style.overflow = "hidden";
      } else {
        document.body.style.overflow = "auto";
      }

      return () => {
        document.body.style.overflow = "auto";
      };
    }, [isOpen]);

    if (!isOpen) {
      return null;
    }

    return (
      <div
        className="fixed inset-0 p-8 bg-black bg-opacity-50 w-screen h-screen top-0 left-0 right-0 bottom-0 overflow-hidden flex justify-center items-center z-30"
        onClick={onClose}
      >
        <WrappedComponent {...(rest as T)} onClose={onClose} />
      </div>
    );
  };

  ModalComponent.displayName = `withModal(${
    WrappedComponent.displayName || WrappedComponent.name || "Component"
  })`;

  return ModalComponent;
}
